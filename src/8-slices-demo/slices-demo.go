package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers = []int{1, 3, 2, 8, 5, 10, 9, 25, 11}
	fmt.Println(numbers)
	fmt.Println("Length of Numbers:", len(numbers))
	fmt.Println("Capacity of Numbers:", cap(numbers))

	numbers = append(numbers[1:len(numbers)])

	fmt.Println(numbers)
	fmt.Println("Length of Numbers:", len(numbers))
	fmt.Println("Capacity of Numbers:", cap(numbers))

	numbers = append(numbers[1:])

	fmt.Println(numbers)
	fmt.Println("Length of Numbers:", len(numbers))
	fmt.Println("Capacity of Numbers:", cap(numbers))

	numbers = append(numbers[0 : len(numbers)-1])

	fmt.Println(numbers)
	fmt.Println("Length of Numbers:", len(numbers))
	fmt.Println("Capacity of Numbers:", cap(numbers))

	numbers = append(numbers[:len(numbers)-1])

	fmt.Println(numbers)
	fmt.Println("Length of Numbers:", len(numbers))
	fmt.Println("Capacity of Numbers:", cap(numbers))

	moreNumbers := make([]int, 5, 5)
	moreNumbers[0] = 57
	moreNumbers[1] = 98
	moreNumbers[2] = 9
	moreNumbers[3] = 8
	moreNumbers[4] = 74

	fmt.Println(moreNumbers)

	moreNumbers = append(moreNumbers, 23)
	fmt.Println(moreNumbers)
	fmt.Println(cap(moreNumbers))

	sort.Ints(moreNumbers)
	fmt.Println(moreNumbers)
}
