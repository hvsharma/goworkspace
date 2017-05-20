package main

import (
	"fmt"
)

//Animal :Useless
type Animal interface {
	Speak() string
}

//Dog :useless
type Dog struct {
	name string
}

//Cat :useless
type Cat struct {
	name string
}

//Cow :useless
type Cow struct {
	name string
}

//Speak : Useless
func (c Cat) Speak() string {
	return "Meow!"
}

//Speak :Does nothing
func (d Dog) Speak() string {
	return "Wow!"
}

//Speak :Does nothing
func (c Cow) Speak() string {
	return "Meme!"
}

func main() {
	dog := Dog{"Poodle"}
	fmt.Println("dog.speak():", dog.Speak())

	//Casting to Animal!  --> If that works, it means Dog implements interface Animal
	a := Animal(dog)
	fmt.Println("a.Speak():", a.Speak())

	cat := Cat{"Nancy"}
	fmt.Println("cat.speak():", cat.Speak())

	//Casting to Animal!  --> If that works, it means Cat implements interface Animal
	b := Animal(cat)
	fmt.Println("b.Speak():", b.Speak())

	cow := Cow{"Shyaama"}
	fmt.Println("cow.speak():", cow.Speak())

	//Casting to Animal!  --> If that works, it means Dog implements interface Animal
	c := Animal(cow)
	fmt.Println("c.Speak():", c.Speak())

	animalSplice := []Animal{dog, cat, cow}

	for i, animal := range animalSplice {
		fmt.Printf("\n%v : %v\n", i, animal)
	}
}
