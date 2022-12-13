package main

import (
	"fmt"
	"net/http"
)

var chance []int
var Nom string
var level string

func page1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/page1.html")
}

func main() {
	http.HandleFunc("/page1", page1)

	portNumber := ":4000"

	fmt.Println("C'est sur" + portNumber)
	http.ListenAndServe(portNumber, nil)
}
