package main

import (
	"encoding/csv"
	"errors"
	"fmt"
)

func checkTaskAsComplete(taskID string) (err error) {
	// Get the index.
	index, err = getIndexByTaskID(taskID)
	if err != nil {
		log.Println(err)
		return
	}

	// Lock the tasks.
	accessTasks.Lock()
	defer accessTasks.Unlock()

	// Set the task to checked.
	allTasks.checked = true
}
