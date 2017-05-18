package main

import "fmt"

//Animal : Useless
type Animal struct {
	name  string
	food  string
	sound string
}

func main() {
	dog := Animal{"Dog", "Omnivore", "Woof!"}
	fmt.Println("Animal:", dog)
	fmt.Printf("Animal: %+v\n", dog)
	fmt.Printf("Name: %v\nFood: %v\nSound: %v\n\n", dog.name, dog.food, dog.speak())
	fmt.Println("Speak by value: ", dog.speakTwoTimesByValue())
	fmt.Println("Speak by value: ", dog.speakTwoTimesByValue())
	fmt.Println("Speak : ", dog.speak())
	fmt.Println("Speak by reference: ", dog.speakTwoTimesByReference())
	fmt.Println("Speak by reference: ", dog.speakTwoTimesByReference())
	fmt.Println("Speak by value: ", dog.speakTwoTimesByValue())
	fmt.Println("Speak by value: ", dog.speakTwoTimesByValue())
	fmt.Println("Speak : ", dog.speak())
}

func (a Animal) speak() string {
	return a.sound
}

func (a Animal) speakTwoTimesByValue() string {
	a.sound = fmt.Sprintf("%v %v", a.sound, a.sound)
	return a.sound
}

func (a *Animal) speakTwoTimesByReference() string {
	a.sound = fmt.Sprintf("%v %v", a.sound, a.sound)
	return a.sound
}
