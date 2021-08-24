package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
		
	//Read in the float - FYI - No error checking on any of this
	fmt.Printf("Hello!  Please enter a float:")

	var floatInput float64;
	fmt.Scanln(&floatInput)
	stringInput := strconv.FormatFloat(floatInput, 'f', -1, 64)

	//Trunc the value and print it out
	fmt.Println("You just entered: " + stringInput + "   We will now trunc it....")
	var truncVal = math.Trunc(floatInput)
	truncString := strconv.FormatFloat(truncVal, 'f', -1, 64)
	fmt.Println("Truncated value: ", truncString)


// 	package main
// import "fmt"
// func main() {
// 	var inputNo float64
// 	fmt.Print("Enter floating point number :: ")
// 	fmt.Scan(&inputNo)
// 	fmt.Printf("Truncated version of '%v' is '%v'.\n", inputNo, int64(inputNo))
// }
	
}