package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func DeleteTask(respWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var remove Remove
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &remove)
	if err != nil {
		fmt.Println(err)
	}

	removeTaskByID(remove.ID)

}
