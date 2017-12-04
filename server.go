package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveData)
	http.HandleFunc("/responsive/", responsiveData)
	http.ListenAndServe(":8080", nil)
}

func serveData(res http.ResponseWriter, req *http.Request) {
	response := "Hello World, from server"
	fmt.Fprintln(res, response)
}

func responsiveData(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/responsive/"):]
	response := "Hello " + name + ", from server"
	fmt.Fprintln(res, response)
}