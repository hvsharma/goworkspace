package main

import "fmt"

func main() {
	doSomething()
	sum := addValues(21, 23)
	fmt.Println("Sum of 2 values:", sum)
	sum = addAllValues(12, 14, 34, 235, 76)
	fmt.Println("Sum of all values:", sum)
}

func doSomething() {
	fmt.Println("Do Something!")
}

//Can also write as -: addValues(x, y int) since both variables have same datatypes
func addValues(x int, y int) int {
	return x + y
}

//Type of variable "values will be a slice!
func addAllValues(values ...int) int {
	fmt.Printf("Type of variable \"values\" is: %T\n", values)
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
