package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var slice []int = make([]int, 3)
	var inputStringOfInts string

	for !isExitValue(inputStringOfInts) {

		fmt.Printf("Please enter Integers (Press x to exit): ")
		fmt.Scanln(&inputStringOfInts) 

		if(isExitValue(inputStringOfInts)) {
			fmt.Println("Exiting...")	
			break;
		}

		intVar, err := strconv.Atoi(inputStringOfInts)
		//fmt.Println(intVar, err, reflect.TypeOf(intVar))
		if(err != nil) {
				fmt.Println("Not an Integer!")
				continue
		} else {
			slice = append(slice, intVar)
			sort.Ints(slice[:])
			fmt.Println(slice)
		}
	}
}

func isExitValue(inputString string) bool {
	return (inputString == "x" || inputString == "X") 
}
