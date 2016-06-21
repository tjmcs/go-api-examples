package main

import (
	"log"
)

func checkTaskAsComplete(taskID int) (err error) {
	// Get the index.
	index, err := getIndexByTaskID(taskID)
	if err != nil {
		log.Println(err)
		return
	}

	// Lock the tasks.
	accessTasks.Lock()
	defer accessTasks.Unlock()

	// Set the task to checked.
	allTasks[index].Checked = true
	return
}
