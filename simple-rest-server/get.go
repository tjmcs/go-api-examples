package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ListTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(respWriter, "Show Details for ToDo")
}

func SearchTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(respWriter, "Show Details for ToDo\n")
}
