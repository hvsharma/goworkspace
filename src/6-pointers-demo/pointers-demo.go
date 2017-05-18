package main

import (
	"fmt"
)

func main() {
	var pointer1 *int

	if pointer1 != nil {
		fmt.Println("Pointer1 is:", *pointer1)
	} else {
		fmt.Println("Pointer1 is nil")
	}

	value := 23
	pointer1 = &value

	fmt.Println("Value is:", value)
	fmt.Println("&Value is:", &value)

	if pointer1 != nil {
		fmt.Println("*Pointer1 is:", *pointer1)
		fmt.Println("Pointer1 is:", pointer1)
	} else {
		fmt.Println("Pointer1 is nil")
	}

	*pointer1 = *pointer1 + 30

	fmt.Println("*Pointer1 is:", *pointer1)
	fmt.Println("Value is:", value)

	pointer2 := &pointer1

	fmt.Println("*Pointer2 is:", *pointer2)
	fmt.Println("Pointer2 is:", pointer2)

	pointer3 := &*pointer1

	fmt.Println("*Pointer3 is:", *pointer3)
	fmt.Println("Pointer3 is:", pointer3)

}
