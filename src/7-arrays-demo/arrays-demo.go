package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)

	var numbers = [5]int{1, 3, 2, 8, 5}
	fmt.Println(numbers)
	fmt.Println("Length of Colors:", len(colors))
	fmt.Println("Length of Numbers:", len(numbers))
}
