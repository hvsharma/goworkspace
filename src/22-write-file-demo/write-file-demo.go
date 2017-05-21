package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is my first ever Go File...!"

	file, err := os.Create("./fromWriteString.txt")
	checkError(err)
	defer file.Close()
	fmt.Println("Writing file...")
	ln, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("File written with io.WriteString()... Length:", ln)

	bytes := []byte(content)
	err = ioutil.WriteFile("./fromBytes.txt", bytes, 0755)
	checkError(err)
	fmt.Println("File written with ioutil.WriteFile()")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
