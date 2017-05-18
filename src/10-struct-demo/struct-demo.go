package main

import (
	"fmt"
)

type Animal struct {
	name string
	food string
}

func main() {
	dog := Animal{"Dog", "Omnivore"}
	fmt.Println("Animal:", dog)
	fmt.Printf("Animal: %+v\n", dog)
	fmt.Printf("Name: %v\nFood: %v\n", dog.name, dog.food)
}
