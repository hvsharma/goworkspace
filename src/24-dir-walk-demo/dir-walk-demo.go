package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("Processing Directory:", root)

	//Note that we are passing a function processPath just like a type...!
	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("Errors:", err)
	}

}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
	}
	return nil
}
