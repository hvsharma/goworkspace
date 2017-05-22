package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {

	serviceURL := "https://api.github.com/users/hvsharma"
	fmt.Println("Trying to retrieve data from:", serviceURL)

	client := getHTTPClient()

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
			result := string(bytes)
			fmt.Println("Response from url:", result)
		}
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
