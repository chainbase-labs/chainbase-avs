package coordinator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type FlinkClient struct {
	FlinkGatewayHttpUrl string
	OssAccessKeyId      string
	OssAccessKeySecret  string
	Client              *http.Client
}

// NewFlinkClient creates a new FlinkClient
func NewFlinkClient(flinkGatewayHttpUrl, ossAccessKeyId, ossAccessKeySecret string) *FlinkClient {
	return &FlinkClient{
		FlinkGatewayHttpUrl: flinkGatewayHttpUrl,
		OssAccessKeyId:      ossAccessKeyId,
		OssAccessKeySecret:  ossAccessKeySecret,
		Client:              &http.Client{},
	}
}

// GetChainLatestBlockHeight gets the latest block height of a chain
func (f *FlinkClient) GetChainLatestBlockHeight(chain string) (int64, error) {
	// Create session
	sessionHandle, err := f.createSession()
	if err != nil {
		return 0, errors.Wrap(err, "failed to create session")
	}

	// Create catalog and set execution mode
	err = f.setupCatalogAndExecutionMode(sessionHandle)
	if err != nil {
		return 0, errors.Wrap(err, "failed to setup catalog and execution mode")
	}

	// Define SQL query
	sqlQuery := "SELECT block_number FROM " + chain + ".blocks ORDER BY block_number DESC LIMIT 1;"

	// Execute SQL query
	operationHandle, err := f.executeSQL(sessionHandle, sqlQuery)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute SQL query")
	}

	// Fetch results
	latestBlockHeight, err := f.fetchResults(sessionHandle, operationHandle)
	if err != nil {
		return 0, errors.Wrap(err, "failed to fetch results")
	}

	return latestBlockHeight, nil
}

// createSession creates a session
func (f *FlinkClient) createSession() (string, error) {
	// Define Flink Gateway API endpoint
	url := f.FlinkGatewayHttpUrl + "/v1/sessions"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to create HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.Client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to send HTTP request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(err, "failed to create session, status code: %d", resp.StatusCode)
	}

	var sessionRespData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&sessionRespData)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse response")
	}

	sessionHandle, ok := sessionRespData["sessionHandle"].(string)
	if !ok {
		return "", errors.New("failed to get session handle from response")
	}

	return sessionHandle, nil
}

// setupCatalogAndExecutionMode sets up the catalog and execution mode
func (f *FlinkClient) setupCatalogAndExecutionMode(sessionHandle string) error {
	// Define the Flink SQL query to create the catalog
	setupQueries := []string{
		"CREATE CATALOG paimon WITH (" +
			"'type' = 'paimon'," +
			"'warehouse' = 'oss://network-testnet/warehouse'," +
			"'fs.oss.endpoint' = 'network-testnet.chainbasehq.com'," +
			"'fs.oss.accessKeyId' = '" + f.OssAccessKeyId + "'," +
			"'fs.oss.accessKeySecret' = '" + f.OssAccessKeySecret + "'," +
			"'table-default.merge-engine' = 'deduplicate'," +
			"'table-default.changelog-producer' = 'input'," +
			"'table-default.metastore.partitioned-table' = 'false'," +
			"'table-default.lookup.cache-file-retention' = '1 h'," +
			"'table-default.lookup.cache-max-memory-size' = '256 mb'," +
			"'table-default.lookup.cache-max-disk-size' = '10 gb'," +
			"'table-default.log.scan.remove-normalize' = 'true'," +
			"'table-default.changelog-producer.row-deduplicate' = 'false'," +
			"'table-default.consumer.expiration-time' = '24 h'," +
			"'table-default.streaming-read-mode' = 'file'," +
			"'table-default.orc.bloom.filter.fpp' = '0.00001'," +
			"'table-default.scan.plan-sort-partition' = 'true'," +
			"'table-default.snapshot.expire.limit' = '10000'," +
			"'table-default.snapshot.num-retained.max' = '2000'" +
			");",
		"USE CATALOG paimon;",
		"SET execution.runtime-mode = 'batch';",
		"SET execution.parallelism.default = 1;",
		"SET parallelism.default = 1;",
	}

	for _, query := range setupQueries {
		_, err := f.executeSQL(sessionHandle, query)
		if err != nil {
			return errors.Wrap(err, "failed to execute setup query")
		}
	}

	return nil
}

// executeSQL executes the SQL query
func (f *FlinkClient) executeSQL(sessionHandle, sqlQuery string) (string, error) {
	// Create the request body for the SQL query
	sqlBody := map[string]interface{}{
		"statement": sqlQuery,
	}
	sqlBodyBytes, err := json.Marshal(sqlBody)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal SQL body")
	}

	// Create and send the HTTP request to execute the SQL query
	sqlAPIURL := f.FlinkGatewayHttpUrl + "/v1/sessions/" + sessionHandle + "/statements"
	req, err := http.NewRequest("POST", sqlAPIURL, bytes.NewBuffer(sqlBodyBytes))
	if err != nil {
		return "", errors.Wrap(err, "failed to create HTTP request for SQL query")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.Client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to send HTTP request for SQL query")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrapf(err, "failed to execute SQL query, status code: %d", resp.StatusCode)
	}

	// Parse the response to get the operation handle
	var sqlRespData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&sqlRespData)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse SQL response")
	}
	operationHandle, ok := sqlRespData["operationHandle"].(string)
	if !ok {
		return "", errors.New("failed to get operation handle from response")
	}

	return operationHandle, nil
}

// fetchResults fetches the results from the Flink SQL query
func (f *FlinkClient) fetchResults(sessionHandle, operationHandle string) (int64, error) {
	resultAPIURL := f.FlinkGatewayHttpUrl + "/v1/sessions/" + sessionHandle + "/operations/" + operationHandle + "/result/0"

	for i := 0; i < 1000; i++ {
		resultRespData, err := f.fetchResultData(resultAPIURL)
		if err != nil {
			time.Sleep(30 * time.Second)
			continue
		}

		results, ok := resultRespData["results"].(map[string]interface{})
		if !ok {
			if nextResultURI, ok := resultRespData["nextResultUri"].(string); ok && nextResultURI != "" {
				resultAPIURL = f.FlinkGatewayHttpUrl + nextResultURI
				time.Sleep(10 * time.Second)
				continue
			}
			return 0, errors.New("failed to get results from response")
		}

		data, ok := results["data"].([]interface{})
		if !ok || len(data) == 0 {
			if nextResultURI, ok := resultRespData["nextResultUri"].(string); ok && nextResultURI != "" {
				resultAPIURL = f.FlinkGatewayHttpUrl + nextResultURI
				time.Sleep(10 * time.Second)
				continue
			}
			return 0, errors.New("data is empty and no next result URI")
		}

		row, ok := data[0].(map[string]interface{})
		if !ok {
			return 0, errors.New("failed to get row from data")
		}

		fields, ok := row["fields"].([]interface{})
		if !ok || len(fields) == 0 {
			return 0, errors.New("failed to get fields from row")
		}

		value, ok := fields[0].(json.Number)
		if !ok {
			return 0, errors.New("failed to convert value to json.Number")
		}

		latestBlockNumber, err := value.Int64()
		if err != nil {
			return 0, errors.Wrap(err, "SQL result is not a valid bigint")
		}

		return latestBlockNumber, nil
	}

	return 0, errors.New("failed to retrieve results after 1000 attempts")
}

// fetchResultData fetches the result data from the Flink SQL query
func (f *FlinkClient) fetchResultData(resultAPIURL string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", resultAPIURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HTTP request for fetching results")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send HTTP request for fetching results")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "failed to fetch results, status code: %d", resp.StatusCode)
	}

	var resultRespData map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()
	err = decoder.Decode(&resultRespData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse result response")
	}

	return resultRespData, nil
}
