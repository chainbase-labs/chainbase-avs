package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestRunFlinkBatchSQL(t *testing.T) {
	// Define the Flink SQL Gateway API endpoint
	flinkSQLGatewayURL := "http://localhost:8083/v1/sessions"

	// Create and send the HTTP request to create a session
	req, err := http.NewRequest("POST", flinkSQLGatewayURL, nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request for session: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send HTTP request for session: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed to create session, status code: %d", resp.StatusCode)
	}

	// Parse the response to get the session handle
	var sessionRespData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&sessionRespData)
	if err != nil {
		t.Fatalf("Failed to parse session response: %v", err)
	}
	sessionHandle, ok := sessionRespData["sessionHandle"].(string)
	if !ok {
		t.Fatalf("Failed to get session handle from response")
	}

	// Get OSS access key ID and secret from environment variables
	ossAccessKeyID := os.Getenv("OSS_ACCESS_KEY_ID")
	ossAccessKeySecret := os.Getenv("OSS_ACCESS_KEY_SECRET")

	if ossAccessKeyID == "" || ossAccessKeySecret == "" {
		t.Fatal("OSS_ACCESS_KEY_ID or OSS_ACCESS_KEY_SECRET environment variable is not set")
	}

	// Define the Flink SQL query to create the catalog
	createCatalogQuery := fmt.Sprintf(`CREATE CATALOG paimon WITH (
		'type' = 'paimon',
		'warehouse' = 'oss://network-testnet/warehouse',
		'fs.oss.endpoint' = 'network-testnet.chainbasehq.com',
		'fs.oss.accessKeyId' = '%s',
		'fs.oss.accessKeySecret' = '%s',
		'table-default.merge-engine' = 'deduplicate',
		'table-default.changelog-producer' = 'input',
		'table-default.metastore.partitioned-table' = 'false',
		'table-default.lookup.cache-file-retention' = '1 h',
		'table-default.lookup.cache-max-memory-size' = '256 mb',
		'table-default.lookup.cache-max-disk-size' = '10 gb',
		'table-default.log.scan.remove-normalize' = 'true',
		'table-default.changelog-producer.row-deduplicate' = 'false',
		'table-default.consumer.expiration-time' = '24 h',
		'table-default.streaming-read-mode' = 'file',
		'table-default.orc.bloom.filter.fpp' = '0.00001',
		'table-default.scan.plan-sort-partition' = 'true',
		'table-default.snapshot.expire.limit' = '10000',
		'table-default.snapshot.num-retained.max' = '2000'
	);`, ossAccessKeyID, ossAccessKeySecret)

	useCatalogQuery := `USE CATALOG paimon;`

	setExecutionModeQuery := `SET execution.runtime-mode = 'batch';`

	setParallelismQuery := `SET execution.parallelism.default = 2;`

	// Execute the create catalog query
	t.Log("Executing create catalog query")
	_, err = executeSQL(t, client, sessionHandle, createCatalogQuery)
	if err != nil {
		t.Fatalf("Failed to create catalog: %v", err)
	}

	// Use the catalog
	t.Log("Using the catalog")
	_, err = executeSQL(t, client, sessionHandle, useCatalogQuery)
	if err != nil {
		t.Fatalf("Failed to use catalog: %v", err)
	}

	// Set the execution mode to batch
	t.Log("Setting execution mode to batch")
	_, err = executeSQL(t, client, sessionHandle, setExecutionModeQuery)
	if err != nil {
		t.Fatalf("Failed to set execution mode to batch: %v", err)
	}

	// Set the parallelism to 8
	t.Log("Setting parallelism to 8")
	_, err = executeSQL(t, client, sessionHandle, setParallelismQuery)
	if err != nil {
		t.Fatalf("Failed to set parallelism to 8: %v", err)
	}

	// Define the Flink SQL query
	sqlQuery := `SELECT block_number FROM zkevm.blocks where block_number > 100000 order by block_number desc limit 1;`

	// Execute the SQL query
	t.Log("Executing SQL query")
	operationHandle, err := executeSQL(t, client, sessionHandle, sqlQuery)
	if err != nil {
		t.Fatalf("Failed to execute SQL query: %v", err)
	}

	// Fetch the results using the session handle and operation handle
	resultAPIURL := fmt.Sprintf("http://localhost:8083/v1/sessions/%s/operations/%s/result/0", sessionHandle, operationHandle)
	var resultRespData map[string]interface{}

	for i := 0; i < 1000; i++ {
		t.Logf("Fetching results, attempt %d", i+1)
		resultRespData, err = fetchResults(t, client, resultAPIURL)
		fmt.Println("get resultRespData", resultRespData)

		if err != nil {
			t.Logf("Attempt %d: Failed to fetch results, retrying in 30 seconds... Error: %v", i+1, err)
			time.Sleep(30 * time.Second)
			continue
		}

		// Check if the 'data' field in 'results' is empty
		results, ok := resultRespData["results"].(map[string]interface{})
		if !ok {
			// Check if there's a next result URI
			if nextResultURI, ok := resultRespData["nextResultUri"].(string); ok && nextResultURI != "" {
				resultAPIURL = "http://localhost:8083" + nextResultURI
				time.Sleep(10 * time.Second)
				continue
			}
			t.Fatalf("Failed to get results from response")
		}

		data, ok := results["data"].([]interface{})
		if !ok || len(data) == 0 {
			fmt.Println("resultRespData", resultRespData)
			// Check if there's a next result URI
			if nextResultURI, ok := resultRespData["nextResultUri"].(string); ok && nextResultURI != "" {
				resultAPIURL = "http://localhost:8083" + nextResultURI
				time.Sleep(10 * time.Second)
				continue
			} else {
				t.Fatalf("Failed to get data from results or data is empty and no nextResultUri")
			}
		} else {
			// If data is not empty, proceed with assertions
			row, ok := data[0].(map[string]interface{})
			if !ok {
				t.Fatalf("Failed to get row from data")
			}

			fields, ok := row["fields"].([]interface{})
			if !ok || len(fields) == 0 {
				t.Fatalf("Failed to get fields from row")
			}

			value, ok := fields[0].(json.Number)
			if !ok {
				t.Fatalf("Failed to get value as json.Number")
			}

			latestBlockNumber, err := value.Int64()
			if err != nil {
				t.Fatalf("SQL result is not a valid bigint: %v", err)
			}

			t.Logf("Successfully retrieved and verified the result: %s", value.String())
			fmt.Printf("Latest block number: %d\n", latestBlockNumber)
			break
		}

		time.Sleep(30 * time.Second)
	}

	fmt.Println("resultRespData", resultRespData)
	t.Fatal("Failed to retrieve results after 10 attempts")
}

func fetchResults(t *testing.T, client *http.Client, resultAPIURL string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", resultAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request for fetching results: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request for fetching results: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch results, status code: %d", resp.StatusCode)
	}

	var resultRespData map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()
	err = decoder.Decode(&resultRespData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result response: %v", err)
	}

	return resultRespData, nil
}

func executeSQL(t *testing.T, client *http.Client, sessionHandle, sqlQuery string) (string, error) {
	// Create the request body for the SQL query
	sqlBody := map[string]interface{}{
		"statement": sqlQuery,
	}
	sqlBodyBytes, err := json.Marshal(sqlBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal SQL body: %v", err)
	}

	// Create and send the HTTP request to execute the SQL query
	sqlAPIURL := fmt.Sprintf("http://localhost:8083/v1/sessions/%s/statements", sessionHandle)
	req, err := http.NewRequest("POST", sqlAPIURL, bytes.NewBuffer(sqlBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request for SQL query: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request for SQL query: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to execute SQL query, status code: %d", resp.StatusCode)
	}

	// Parse the response to get the operation handle
	var sqlRespData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&sqlRespData)
	if err != nil {
		return "", fmt.Errorf("failed to parse SQL response: %v", err)
	}
	operationHandle, ok := sqlRespData["operationHandle"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get operation handle from response")
	}

	return operationHandle, nil
}
