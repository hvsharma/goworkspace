package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i1, i2, i3 := 21, 30, 42
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum =", intSum)

	f1, f2, f3 := 11.1, 13.2, 7.5
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum = ", floatSum)

	var bf1, bf2, bf3, bigFSum big.Float
	bf1.SetFloat64(11.1)
	bf2.SetFloat64(13.2)
	bf3.SetFloat64(7.5)
	bigFSum.Add(&bf1, &bf2).Add(&bigFSum, &bf3)
	fmt.Printf("BigFSum = %.10g", &bigFSum)
	fmt.Println()
	circleRadius := 10.2
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference = %.2f", circumference)
}
