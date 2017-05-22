package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/*
	Made Two struct types to parse the response from the github API
*/

//User : Stores user data
type User struct {
	Login string
	Id    int
	Url   string
}

//UsersWrapper : Stores user data
type UsersWrapper struct {
	Total_count        int
	Incomplete_results bool
	Items              []User
}

func main() {
	serviceURL := "https://api.github.com/search/users?q=hvs"
	client := getHTTPClient()
	content := getContentFromServer(serviceURL, client)
	users := getUsersFromJSON(content)
	if users != nil {
		for _, user := range users {
			fmt.Printf("\nUsername:\t%v\nUser ID:\t%v\nURL:\t\t%v\n", user.Login, user.Id, user.Url)
			fmt.Println()
		}
	} else {
		fmt.Println("No data available for this query...!")
	}

}

func getContentFromServer(serviceURL string, client *http.Client) string {
	var result string
	fmt.Println("Trying to retrieve data from:", serviceURL)
	response, err := client.Get(serviceURL)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nResponse Type: %T\n", response)
		defer response.Body.Close()
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			result = string(bytes)
			if len(strings.TrimSpace(result)) > 0 {
				fmt.Println("Response from url retrieved. Trying to parse Response...")
			} else {
				fmt.Println("Empty Response...")
			}
		}
	}
	return result
}

func getUsersFromJSON(data string) []User {
	var usersWrapper UsersWrapper
	decoder := json.NewDecoder(strings.NewReader(data))

	for decoder.More() {
		err := decoder.Decode(&usersWrapper)
		checkError(err)
		fmt.Println("Total Items Present:", usersWrapper.Total_count)
		if usersWrapper.Total_count > 0 {
			return usersWrapper.Items
		}
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	getHTTPClient : Added to include the proxy from my own office
	It's best practice to include our own client instead of directly using http.get because
	of the timeout. Source : https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
*/
func getHTTPClient() *http.Client {
	var client *http.Client
	hostname, _ := os.Hostname()

	if hostname == "01HW477850" {
		proxyString := "http://www.mc.xerox.com:8000"
		proxyURL, _ := url.Parse(proxyString)
		tr := &http.Transport{
			Proxy:              http.ProxyURL(proxyURL),
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}

		client = &http.Client{
			Transport: tr,
			Timeout:   time.Second * 10}

	} else {
		client = &http.Client{Timeout: time.Second * 10}
	}

	return client
}
