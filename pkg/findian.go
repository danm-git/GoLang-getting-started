package main

import (
	"fmt"
	"strings"
)

func main() {

	// var cadenai string
	// fmt.Println("Plase input string:")
	// fmt.Scan(&cadenai)

	// cadena := strings.ToLower(cadenai)
	// if strings.Index(cadena, "i") == 0 && strings.Index(cadena, "n") == len(cadena)-1 && strings.Index(cadena, "a") > 0 {
	// 	fmt.Println("Found!")
	// } else {
	// 	fmt.Println("Not Found!")
	// }



	//Read in the float - FYI - No error checking on any of this
	fmt.Printf("Please enter a String:")

	var stringInput string;
	fmt.Scanln(&stringInput)
	//fmt.Println("You entered: ", stringInput)

	found := false;
	inputLen := len(stringInput)
	//fmt.Println("inputLen: ", inputLen)


	// First and last need to be a and n so skipping
	var currChar string;

	for i := 0; i < inputLen -1; i++ {
		currChar = strings.ToLower(string(stringInput[i]))
		//fmt.Println("currChar: ", currChar)
		if(currChar == "a") {
			found = true
			break;
		}
	}
	
	//fmt.Println("found: ", found)

	if(found) {
		firstLetter := strings.ToLower(string(stringInput[0]))
		lastLetter := strings.ToLower(string(stringInput[len(stringInput) - 1]))


		isFirstValid := firstLetter == string("i")
		//fmt.Println("firstLetter: ", firstLetter)
		//fmt.Println("isFirstValid: ", isFirstValid)
		
		isLastValid := lastLetter == string("n")
		//fmt.Println("lastLetter: ", lastLetter)
		//fmt.Println("isLastValid: ", isLastValid)

		found = isFirstValid && isLastValid
	}

	if(found) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
		
}

