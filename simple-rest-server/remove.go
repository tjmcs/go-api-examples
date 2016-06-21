package main

import (
	"log"
)

func removeTaskByID(taskID int) (err error) {
	index, err := getIndexByTaskID(taskID)
	if err != nil {
		log.Println(err)
		return
	}
	accessTasks.Lock()
	allTasks = append(allTasks[:index-1], allTasks[index+1:]...)
	accessTasks.Unlock()
	return
}
