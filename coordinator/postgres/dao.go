package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type Operator struct {
	OperatorAddress string
	OperatorID      string
	Socket          string
	Location        string
	CPUCore         uint32
	Memory          uint32
	Status          string
}

type Task struct {
	TaskID         uint32
	TaskDetail     string
	TaskResponse   string
	CreateTaskTx   string
	ResponseTaskTx string
}

type OperatorTask struct {
	OperatorID uint32
	TaskID     uint32
	Status     string
}

func UpsertOperator(db *sql.DB, operator *Operator) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO operator (
            operator_address,
            operator_id,
            socket,
            location,
            cpu_core,
            memory,
            status
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (operator_address) DO UPDATE SET
            operator_id = EXCLUDED.operator_id,
            socket = EXCLUDED.socket,
            location = EXCLUDED.location,
            cpu_core = EXCLUDED.cpu_core,
            memory = EXCLUDED.memory,
            status = EXCLUDED.status,
            updated_at = CURRENT_TIMESTAMP
        RETURNING id`,
		operator.OperatorAddress,
		operator.OperatorID,
		operator.Socket,
		operator.Location,
		operator.CPUCore,
		operator.Memory,
		operator.Status,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to upsert operator: %w", err)
	}
	return id, nil
}

func UpdateOperatorRegisteredAt(db *sql.DB, operatorAddress string, timestamp time.Time) error {
	registeredAt := sql.NullTime{Time: timestamp, Valid: true}
	_, err := db.Exec(`UPDATE operator SET registered_at = $1 WHERE operator_address = $2`, registeredAt, operatorAddress)
	return err
}

func QueryOperatorAddressesNoRegisteredAt(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT operator_address FROM operator WHERE registered_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := make([]string, 0)
	for rows.Next() {
		var address string
		if err := rows.Scan(&address); err != nil {
			continue
		}
		addresses = append(addresses, address)
	}
	return addresses, err
}

func InsertTask(db *sql.DB, task *Task) error {
	_, err := db.Exec(`
        INSERT INTO task (task_id, task_detail, create_task_tx) 
        VALUES ($1, $2, $3)
        ON CONFLICT (task_id) DO NOTHING;`,
		task.TaskID,
		task.TaskDetail,
		task.CreateTaskTx,
	)
	return err
}

func UpdateTaskResponse(db *sql.DB, taskID uint32, taskResponse, responseTaskTx string) error {
	_, err := db.Exec(`UPDATE task SET task_response = $1, response_task_tx = $2 WHERE task_id = $3 AND task_response is NULL`, taskResponse, responseTaskTx, taskID)
	return err
}

func BatchInsertOperatorTasks(db *sql.DB, operatorTasks []*OperatorTask) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
        INSERT INTO operator_task (operator_id, task_id, status) 
        VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, ot := range operatorTasks {
		_, err = stmt.Exec(ot.OperatorID, ot.TaskID, ot.Status)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func CountOperatorTasks(db *sql.DB, taskID uint32) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM operator_task WHERE task_id = $1`

	err := db.QueryRow(query, taskID).Scan(&count)
	return count, err
}

func QueryOperatorIDs(db *sql.DB, operatorAddresses []string) ([]uint32, error) {
	rows, err := db.Query("SELECT id FROM operator WHERE operator_address = ANY($1)", pq.Array(operatorAddresses))
	if err != nil {
		return nil, fmt.Errorf("query operator ids failed: %w", err)
	}
	defer rows.Close()

	var operatorIDs []uint32

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan operator id failed: %w", err)
		}
		operatorIDs = append(operatorIDs, uint32(id))
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate rows failed: %w", err)
	}

	return operatorIDs, nil
}
