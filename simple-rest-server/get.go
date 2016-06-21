package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

func ListTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	output, _ := json.MarshalIndent(allTasks, "", "  ")
    fmt.Fprintln(respWriter, string(output))
}

func SearchTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(respWriter, "Show Details for ToDo\n")
}
