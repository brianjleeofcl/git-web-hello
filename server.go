package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveData)
	http.ListenAndServe(":8080", nil)
}

func serveData(res http.ResponseWriter, req *http.Request) {
	response := "Hello World, from server"
	fmt.Fprintln(res, response)
}