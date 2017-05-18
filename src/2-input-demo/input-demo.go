package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enter a number: ")
	str, _ = reader.ReadString('\n')
	i, err := strconv.ParseInt(strings.TrimSpace(str), 10, 32)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
