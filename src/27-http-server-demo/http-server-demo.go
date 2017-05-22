package main

import (
	"fmt"
	"net/http"
)

//HTTPServer : Empty Struct
type HTTPServer struct{}

func (h HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style=\"font-family:calibri;\">Yo Bro! Sup?</h1>")
}

func main() {
	var h HTTPServer
	err := http.ListenAndServe("localhost:2000", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
