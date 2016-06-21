package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
)

func getIndexByTaskID(taskID string) (index int, err error) {
	// We didn't find the task ID - return an error and set the index to -1...
	err = error.New("GetIndexByTaskID: TaskID " + taskID + " Not Found.")
	index = -1

	accessTasks.Lock()
	defer accessTasks.Unlock()

	// Search for the taskID and return the index.
	for i, line := range allTask {
		if taskID == line.ID {
			index = i
			err = nil
			return
		}
	}
}

type Task struct {
	ID        int
	Checked   bool      `json:"checked"`
	TimeAdded time.Time `json:"time_added"`
	Deadline  time.Time `json:"deadline"`
	Task      string    `json:"task"`
}

var (
	allTasks    []Task
	accessTasks = &sync.Mutex{}
)

const (
	timeFormat = time.RFC3339
)

func main() {

	// Loading the csv file into the RAM
	csvfile, err := os.Open("tasks.csv")
	panic(err)
	rawCSVdata, err := csv.NewReader(csvfile).ReadAll()
	panic(err)

	for i, each := range rawCSVdata {
		timeAdded, err := time.Parse(timeFormat, each[2])
		panic(err)
		deadline, err := time.Parse(timeFormat, each[3])
		panic(err)
		status := false
		if each[1] == "true" {
			status = true
		}
		allTasks = append(allTasks, Task{ID: i, Task: each[4], TimeAdded: timeAdded, Deadline: deadline, Checked: status})
	}

	// Autosave every minutes
	csvfile.Close()
	go func() {
		for {
			saveCSV()
			time.Sleep(1 * time.Minute)
		}
	}()

	// Start the API
	router := httprouter.New()
	router.GET("/search", SearchTask)
	router.GET("/list", ListTask)
	router.POST("/add", AddTask)
	router.DELETE("/delete", DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func saveCSV() {
	myString := ""
	accessTasks.Lock()
	for _, each := range allTasks {
		myString += fmt.Sprintf("%v,%v,%v,%v,\"%v\"\n", each.ID, each.Checked, each.TimeAdded.Format(timeFormat), each.Deadline.Format(timeFormat), each.Task)
	}
	ioutil.WriteFile("tasks.csv", []byte(myString), 0644)
	accessTasks.Unlock()
}

func panic(err error) {
	if err != nil {
		// Maybe change to log later
		fmt.Println(err)
		os.Exit(1)
	}
}
