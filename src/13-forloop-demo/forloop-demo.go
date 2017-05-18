package main

import "fmt"

func main() {
	colors := make([]string, 7)
	colors[0] = "Violet"
	colors[1] = "Indigo"
	colors[2] = "Blue"
	colors[3] = "Green"
	colors[4] = "Yellow"
	colors[5] = "Orange"
	colors[6] = "Red"

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("\nNow using range in for loop...")
	for i := range colors {
		fmt.Println(colors[i])
	}

	fmt.Println("\nNow using while loop (In GO, There is no dedicated while loop. It's just a variation of for loop)")
	limit := len(colors)
	i := 0

	for i < limit {
		fmt.Println(colors[i])
		i++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 500 {
			goto endOfprogram
		}
	}

endOfprogram:

	fmt.Println("This is the end of the world!")

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 500 {
			break
		}
	}

	fmt.Println("This is the end of break-up!")

	sum = 1
	for sum < 1000 {
		if sum > 500 && sum < 600 {
			sum += 10
			fmt.Println("Sum:", sum)
			continue
		}
		sum += sum
		fmt.Println("Sum:", sum)
	}

	fmt.Println("Value of sum after continue is:", sum)
}
