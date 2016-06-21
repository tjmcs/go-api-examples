package main

import (
	"fmt"
	"net/http"
    "os"
	"encoding/csv"
	"github.com/julienschmidt/httprouter"
)

func ListTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	 csvfile, err := os.Open("tasks.csv")

         if err != nil {
                 fmt.Println(err)
                 return
         }

         defer csvfile.Close()

         reader := csv.NewReader(csvfile)

         reader.FieldsPerRecord = -1 
         
         rawCSVdata, err := reader.ReadAll()

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // sanity check, display to standard output
         for _, each := range rawCSVdata {
             status := "unchecked"
             if each[1] == "true" {  // each[1] will be the status - either true or false
                 status = "checked"
             }
              fmt.Fprintln(respWriter,"ID:", each[0], "Status :", status, "Date Added:", each[2], "Date Due:", each[3], "Task:", each[4])
         }
}

func SearchTask(respWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(respWriter, "Show Details for ToDo\n")
}
