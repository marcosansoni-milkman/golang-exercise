package handler

import (
	"encoding/json"
	"fmt"
	"http/dto"
	"log"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %v \n", r)
	queryString := r.URL.Query()
	name, ok := queryString["name"]

	if !ok {
		e := dto.Error{Code: 100, Message: "Invalid Name"}
		jsonData, _ := json.Marshal(dto.Response{Errors: []dto.Error{e}, Success: false})
		fmt.Printf("Response %v \n", string(jsonData))
		w.Write(jsonData)
		return
	}
	log.Println(name)
	res := dto.ResultWrapper{Message: "Hi " + name[0]}
	//res := []byte(`{"greet": "Hi `+name[0]+`"}`)
	successJson, _ := json.Marshal(dto.Response{Success: true, Result: res})

	log.Printf("Response: %v \n", string(successJson))
	w.Write(successJson)
}
