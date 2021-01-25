package middleware

import (
	"encoding/json"
	"fmt"
	"http/dto"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		//log.Println(r.RequestURI)

		//queryString := r.URL.Query()
		session := r.Header.Get("session")
		//session, ok := queryString["session"]

		if session != "validSession" {
			log.Println("Not a valid session")

			e := dto.Error{Code: 2200, Message: "Invalid session"}
			res := dto.Response{Success: false, Errors: []dto.Error{e}}
			jsonData, _ := json.Marshal(res)

			w.Write(jsonData)
			return
		}

		fmt.Println("Forward")

		log.Println("Valid session")
		//return
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
