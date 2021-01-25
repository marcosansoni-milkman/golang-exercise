package handler

import (
	"fmt"
	"net/http"
	"time"
)

var StartServer time.Time

func init() {
	StartServer = time.Now()
	fmt.Println("init")
}

func IsActive(w http.ResponseWriter, r *http.Request) {
	response := []byte(`{
						"success": 200 
						"result": {
							"startTime" : ` + StartServer.Format(time.RFC3339Nano) + `
						}
					}`)
	fmt.Printf("%v \n", r)
	w.Write(response)
}
