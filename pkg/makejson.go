package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var inputName string
	var inputAddress string

	fmt.Printf("Please enter your name: ")
	fmt.Scanln(&inputName) 

	fmt.Printf("Please enter your address: ")
	fmt.Scanln(&inputAddress) 

	myMap := make(map[string]string)
	myMap["name"] = inputName
	myMap["address"] = inputAddress

	jsonString, err := json.Marshal(myMap)
	
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonString)) 
	}
}


