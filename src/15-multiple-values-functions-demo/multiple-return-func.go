package main

import "fmt"

func main() {
	fN, len := FullName("Neo", "Anderson")
	fmt.Printf("\nFullname: %v, Length: %v\n", fN, len)

	fN2, len2 := FullNameNakedReturn("Harsh", "Sharma")
	fmt.Printf("\nFullname: %v, Length: %v\n", fN2, len2)
}

//FullName :Returns fullname by concatenating first and last name
func FullName(firstName, lastName string) (string, int) {
	fullName := firstName + " " + lastName
	length := len(fullName)
	return fullName, length
}

//FullNameNakedReturn :Returns fullname by concatenating first and last name but here, variable
//names have been mentioned in the return type itself and therefore need not be explicitly returned
func FullNameNakedReturn(firstName, lastName string) (fullName string, length int) {
	fullName = firstName + " " + lastName
	length = len(fullName)
	return
}
