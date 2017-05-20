package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("NotExistingFile.txt")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err)
	}

	fmt.Println(err.Error())

	//To Override the error message...

	myError := errors.New("this is the new customized error")
	fmt.Println(myError)

	attendance := map[string]bool{
		"Charlie": false,
		"Napster": true}

	//ok : will return true if the key exists, otherwise will return false
	//isPresent : will have the value relating to the key
	isPresent, ok := attendance["Charlie"]

	if ok {
		fmt.Println("Charlie is Present? :", isPresent)
	} else {
		fmt.Println("No information for Charlie...!")
	}

	isPresent, ok = attendance["Neo"]

	if ok {
		fmt.Println("Neo is Present? :", isPresent)
	} else {
		fmt.Println("No information for Neo...!")
	}
}
