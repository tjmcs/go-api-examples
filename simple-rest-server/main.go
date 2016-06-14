package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(respWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintln(respWriter, "ToDo RESTful App")
}

func TodoIndex(respWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintln(respWriter, "Return ToDo List")
}

func TodoShow(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(respWriter, "Show Details for ToDo '%s'\n", params.ByName("todoId"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
