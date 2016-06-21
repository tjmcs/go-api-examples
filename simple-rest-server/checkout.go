package main

import (
	"encoding/csv"
	"errors"
	"fmt"
)

func getIndexByTaskID(taskID string, csv [][]string) (index int, err error) {
	for i, line := range csv {
		if taskID == line[0] {
			index = i
			return
		}
	}
	err = error.New("GetIndexByTaskID: TaskID " + taskID + " Not Found.")
	index = -1
}

func checkTaskAsComplete(taskID string, csv [][]string) (err error) {
	index, err = getIndexByTaskID(taskID, csv)
	if err != nil {
		log.Println(err)
		return err
	}
	csv[index][1] = 1
}
