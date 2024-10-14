package node

import (
	"fmt"
	"time"

	"github.com/chainbase-labs/chainbase-avs/core"
)

// TestManuscriptNodeTask run a test task on manuscript node
func (n *ManuscriptNode) TestManuscriptNodeTask() error {
	deadline := time.Now().Add(time.Hour).Unix()
	taskDetails := fmt.Sprintf("v1;zkevm;block;merkle;start:3183646;end:3183746;difficulty:10;deadline:%d", deadline)

	parsedTaskDetails, err := core.ParseTaskDetails(taskDetails)
	if err != nil {
		n.logger.Error("Failed to parse task details", "err", err)
		return err
	}

	if err = n.ExecuteTask(0, parsedTaskDetails); err != nil {
		n.logger.Error("Failed to execute task", "err", err)
		return err
	}

	if err = n.WaitTaskCompletion(0, parsedTaskDetails); err != nil {
		n.logger.Error("Error wait task completion", "err", err)
		n.CancelTaskJob(0)
		return err
	}

	_, err = n.QueryTaskResponse(0)
	if err != nil {
		n.logger.Error("Error query task response", "err", err)
		n.CancelTaskJob(0)
		return err
	}

	n.CancelTaskJob(0)

	return nil
}
