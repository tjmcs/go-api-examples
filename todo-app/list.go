package main

 import (
         "encoding/csv"
         "fmt"
         "os"
 )

 func main() {

         csvfile, err := os.Open("somecsvfile.csv")

         if err != nil {
                 fmt.Println(err)
                 return
         }

         defer csvfile.Close()

         reader := csv.NewReader(csvfile)

         reader.FieldsPerRecord = -1 // see the Reader struct information below

         rawCSVdata, err := reader.ReadAll()

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // sanity check, display to standard output
         for _, each := range rawCSVdata {
             status := "unchecked"
             if(each[1]){  // each[1] will be the status - either true or false
                 status = "checked"
             }
                 fmt.Fprintln("ID: %s , Status : %s , Date Added: %s , Date Due: %s , Task: %s", each[0], status, each[2], each[3], each[4])
         }
 }