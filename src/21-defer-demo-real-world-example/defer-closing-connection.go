package main

import (
	"fmt"
)

var isConnected = false

func main() {
	fmt.Println("Connection status: ", isConnected)
	doSomething()
	fmt.Println("Connection status: ", isConnected)
}

func doSomething() {
	connect()
	fmt.Println("Deferring disconnect...")
	defer disconnect()
	fmt.Println("Connection status: ", isConnected)
	fmt.Println("Doing something...")
}

func connect() {
	fmt.Println("Opening connection...")
	isConnected = true
}

func disconnect() {
	fmt.Println("Disconnecting connection...!")
	isConnected = false
}
