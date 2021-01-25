package handler

import (
	"encoding/json"
	"fmt"
	"http/dao"
	"http/dto"
	"net/http"
	"time"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user dto.User
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&user)
	now := time.Now()
	//user := dto.User{CreatedAt: now, UpdatedAt: now}

	//userTest := dto.User{}
	//
	//test := json.Unmarshal(r.Body, &userTest)
	user.CreatedAt = now
	user.UpdatedAt = now
	//err := jsonDecoder.Decode(&user)

	//fmt.Println("User details")
	//fmt.Println(user.Name)
	//fmt.Println(user.CreatedAt)
	//fmt.Println(user.UpdatedAt)
	//fmt.Println(user.Surname)
	//fmt.Println("Error decoding")
	//fmt.Println(err)
	//fmt.Println("---------")

	if err != nil {
		e := dto.Error{Code: 100, Message: "Invalid User}"}
		jsonData, _ := json.Marshal(dto.Response{Errors: []dto.Error{e}, Success: false})
		fmt.Printf("Response %v \n", string(jsonData))
		w.Write(jsonData)
		return
	}

	u, errInsert := dao.InsertUser(user)
	if errInsert != nil {
		e := dto.Error{Code: 100, Message: "Error while inserting user}"}
		jsonData, _ := json.Marshal(dto.Response{Errors: []dto.Error{e}, Success: false})
		fmt.Printf("Response %v \n", string(jsonData))
		w.Write(jsonData)
		return
	}
	fmt.Println(u)

	//responseByte := []byte(`{
	//					"success": 200
	//					"result": {
	//						"startTime" : ` + StartServer.Format(time.RFC3339Nano) + `
	//					}
	//				}`)

	response, _ := json.Marshal(dto.Response{Success: true, Result: dto.ResultWrapper{Entity: dto.Entity{User: u}}})
	fmt.Printf("%v \n", r)
	w.Write(response)

}
