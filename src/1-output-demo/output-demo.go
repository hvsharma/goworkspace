package main

import (
	"fmt"
	"strings"
)

func main() {
	stringLength, err := fmt.Println(strings.ToUpper("Yo Bro! What's up!"))
	var aNumber = 10
	var isTrue = true
	var str = "Hello World!"
	if err == nil {
		fmt.Println("String length is -:", stringLength)
	} else {
		fmt.Println(err)
	}

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber in float: %.2f\n", float64(aNumber))

	_, err = fmt.Printf("Value of str in numerical format is: %v\n", str)

	if err == nil {
		fmt.Println("String length is -:", stringLength)
	} else {
		fmt.Println(err)
	}
}
