package main

import (
	"basic/arrays"
	"basic/map"
	"basic/struct"
	"fmt"
)

func main() {
	fmt.Println("Array with number")
	res := arrays.NumberArray(5)
	for _, item := range res {
		fmt.Println(item)
	}

	fmt.Println("Array with string")
	stringArrayRes := arrays.StringArray(10)
	for _, item := range stringArrayRes {
		fmt.Println(item)
	}

	fmt.Println("Struct - House")
	h := _struct.House{Price: 100000, NumberRoom: 2}
	fmt.Println(h.Evaluate())
	fmt.Printf("Is two room apartment? %t \n", h.IsTwoRoomApartment())

	fmt.Println("Map")
	mapRes := _map.AgeByName()
	// Iterate over map keys
	for item, _ := range mapRes {
		fmt.Printf("%s with age %d \n", item, mapRes[item])
	}

}
