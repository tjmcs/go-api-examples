package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/search", SearchTask)
	router.GET("/list", ListTask)
	router.POST("/add", AddTask)
	router.DELETE("/delete", DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", router))
}
