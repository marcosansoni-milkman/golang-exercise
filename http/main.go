package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"http/handler"
	"http/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.AuthMiddleware)

	router.HandleFunc("/active", handler.IsActive).Methods(http.MethodGet)
	router.HandleFunc("/greeting", handler.Greeting).Methods(http.MethodGet)
	router.HandleFunc("/user", handler.AddUser).Methods(http.MethodPost)

	err := http.ListenAndServe(":8082", router)

	if err != nil {
		fmt.Println("Error on creation of a connection")
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connection successfully created")
}
