package main

import (
	"fmt"
	"stringutil"
)

func main() {
	fN, len := stringutil.FullName("Neo", "Anderson")
	fmt.Printf("\nFullname: %v, Length: %v\n", fN, len)

	fN2, len2 := stringutil.FullNameNakedReturn("Harsh", "Sharma")
	fmt.Printf("\nFullname: %v, Length: %v\n", fN2, len2)
}
