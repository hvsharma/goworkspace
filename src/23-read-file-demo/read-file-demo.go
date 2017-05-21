package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Attempting to read file...")
	data, err := ioutil.ReadFile("./hello.txt")
	checkError(err)
	fmt.Println()
	fmt.Println("Data from file:", data)
	result := string(data)
	fmt.Println()
	fmt.Println("Data converted into string: ", result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
