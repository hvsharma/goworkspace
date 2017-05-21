package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.github.com/users/hvsharma"
	fmt.Println("Trying to retrieve data from:", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nResponse Type: %T\n", response)
		defer response.Body.Close()
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			result := string(bytes)
			fmt.Println("Response from url:", result)
		}
	}
}
