package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AddTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(respWriter, "Show Details for ToDo '%s'\n", params.ByName("todelete"))
}
