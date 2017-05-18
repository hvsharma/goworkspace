package main

import "fmt"

func main() {
	num := 2

	if num > 0 {
		fmt.Println("Result: ", "Greater than zero!")
	} else if num == 0 {
		fmt.Println("Result: ", "Equal to zero!")
	} else {
		fmt.Println("Result: ", "Less than zero!")
	}

	fmt.Println("With inline variable declaration in If statement -:")

	if num2 := 0; num2 > 0 {
		fmt.Println("Result: ", "Greater than zero!")
	} else if num2 == 0 {
		fmt.Println("Result: ", "Equal to zero!")
	} else {
		fmt.Println("Result: ", "Less than zero!")
	}

	fmt.Println("Num2 will be undefined outside if statement becz it will be out of scope...!")
}
