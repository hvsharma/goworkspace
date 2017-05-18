package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "Implicit type inference!"
	var str2 string = "Explicitly typed string!"

	fmt.Printf("Value and Type of str1 - %v:%T\n", str1, str1)
	fmt.Printf("Value and Type of str2 - %v:%T\n", str2, str2)

	value1 := "hello"
	value2 := "HELLO"

	fmt.Printf("Are Strings \"%v\" and \"%v\" are equal? "+strconv.FormatBool((value1 == value2))+"\n\n", value1, value2)
	fmt.Println("Using EqualFold (Non Case Sensitive)...")
	fmt.Printf("Are Strings \"%v\" and \"%v\" are equal? "+strconv.FormatBool(strings.EqualFold(value1, value2))+"\n\n", value1, value2)
	fmt.Println("Using Contains...")
	fmt.Printf("Does str1 contains \"Exp\" ? " + strconv.FormatBool(strings.Contains(str1, "Exp")) + "\n")
	fmt.Printf("Does str2 contains \"Exp\" ? " + strconv.FormatBool(strings.Contains(str2, "Exp")) + "\n")
}
