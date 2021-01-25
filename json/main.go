package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name     string     `json:"name"`
	Surname  string     `json:"surname"`
	Age      int        `json:"age"`
	Language []Language `json:"language"`
}

type Language struct {
	Name            string `json:"name"`
	MonthExperience int    `json:"monthExperience"`
}

func main() {
	fmt.Println("Hello World")

	jsonData := []byte(`{
		"name": "Marco",
		"surname": "Sansoni",
		"age": 25,
		"language": [
			{
				"name": "JS",
				"monthExperience": 18
			}
		]
	}`)

	if !json.Valid(jsonData) {
		fmt.Println("JSON is not properly formatted")
		os.Exit(1)
	}

	fmt.Println("JSON is formatted properly")

	var p Person
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s is %d years old and knows %d languages \n", p.Name, p.Age, len(p.Language))
}
