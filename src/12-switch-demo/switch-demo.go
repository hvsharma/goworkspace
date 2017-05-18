package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var1 := rand.Intn(10)
	fmt.Println("Value of var1:", var1)

	switch var1 {
	case 1:
		fmt.Println("equal to 1...")
	case 5:
		fmt.Println("equal to 5...")
	case 10:
		fmt.Println("equal to 10...")
	default:
		fmt.Println("Nothing is here...")
	}

	//Note that here I didn't provide the variable on which switch has to be operated!

	switch num2 := rand.Intn(10); {
	case num2 < 5:
		fmt.Println("less than 5...")
	case num2 == 5:
		fmt.Println("equal to 5...")
	case num2 > 5:
		fmt.Println("greater than 5...")
	default:
		fmt.Println("Nothing is here...")
	}
}
