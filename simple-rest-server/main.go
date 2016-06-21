package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getIndexByTaskID(taskID string) (index int, err error) {
	accessTasks.Lock()
	defer accesTasks.Unlock()

	// Search for the taskID and return the index.
	for i, line := range allTask {
		if taskID == line.ID {
			index = i
			return
		}
	}
	// We didn't find the task ID - return an error and set the index to -1...
	err = error.New("GetIndexByTaskID: TaskID " + taskID + " Not Found.")
	index = -1
}

func main() {
	router := httprouter.New()
	router.GET("/search", SearchTask)
	router.GET("/list", ListTask)
	router.POST("/add", AddTask)
	router.DELETE("/delete", DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", router))
}
