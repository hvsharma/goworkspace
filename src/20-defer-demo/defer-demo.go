package main

import (
	"fmt"
)

func main() {

	//Deferred statement runs after all the statements in the function block has been executed.
	//If there are multiple deferred statements, they are added to a stack one by one as they appear in execution sequence in
	//that particular function

	defer fmt.Println("deferred statement")
	fmt.Println("statement 0")
	fmt.Println("statement 1")
	fmt.Println("statement 2")
	fmt.Println("statement 3")
	deferredEvalDemo()
}

func deferredEvalDemo() {
	//Value of any variable or statement inside defer is calculated earlier only!
	fmt.Println()
	x := 100
	defer fmt.Printf("\nDeferred Value of x : %v\n", x)
	x = x + 10
	fmt.Println("Non-deferred statement!: value of x:", x)
}
