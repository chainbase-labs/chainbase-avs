package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

type Operator struct {
	OperatorAddress string
	OperatorID      string
	Socket          string
	Location        string
	CPUCore         uint32
	Memory          uint32
	Status          string
	RegisteredAt    sql.NullTime
}

type Task struct {
	TaskID       uint32
	TaskDetail   string
	TaskResponse string
}

type OperatorTask struct {
	OperatorID int
	TaskID     int
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
            status,
            registered_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (operator_address) DO UPDATE SET
            operator_id = EXCLUDED.operator_id,
            socket = EXCLUDED.socket,
            location = EXCLUDED.location,
            cpu_core = EXCLUDED.cpu_core,
            memory = EXCLUDED.memory,
            status = EXCLUDED.status,
            registered_at = EXCLUDED.registered_at,
            updated_at = CURRENT_TIMESTAMP
        RETURNING id`,
		operator.OperatorAddress,
		operator.OperatorID,
		operator.Socket,
		operator.Location,
		operator.CPUCore,
		operator.Memory,
		operator.Status,
		operator.RegisteredAt,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to upsert operator: %w", err)
	}
	return id, nil
}

func UpdateOperatorStatus(db *sql.DB, operatorAddress string) error {
	_, err := db.Exec(`UPDATE operator SET status = $1 WHERE operator_address = $2`, "inactive", operatorAddress)
	return err
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
        INSERT INTO task (task_id, task_detail, task_response) 
        VALUES ($1, $2, $3)
        ON CONFLICT (task_id) DO NOTHING;`,
		task.TaskID,
		task.TaskDetail,
		task.TaskResponse,
	)
	return err
}

func UpdateTaskResponse(db *sql.DB, taskID uint32, taskResponse string) error {
	_, err := db.Exec(`UPDATE task SET task_response = $1 WHERE task_id = $2 AND task_response is NULL`, taskResponse, taskID)
	return err
}

func BatchInsertOperatorTasks(db *sql.DB, operatorTasks []OperatorTask) error {
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
