package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Remove struct {
	ID int `json:"id"`
}

func AddTask(respWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var task Task
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(respWriter, "Show Details for ToDo '%v'\n", task)

	accessTasks.Lock()
	task.ID = allTasks[len(allTasks)-1].ID + 1

	allTasks = append(allTasks, task)
	accessTasks.Unlock()

	saveCSV()

}
