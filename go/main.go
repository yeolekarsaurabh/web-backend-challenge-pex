package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var previous_fib = 0
var current_fib = 1 // current

func current(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strconv.Itoa(previous_fib))
}

func next(w http.ResponseWriter, r *http.Request) {
	var response = strconv.Itoa(current_fib)
	fmt.Fprint(w, response)
	new_fib := previous_fib + current_fib
	previous_fib = current_fib
	current_fib = new_fib
}

func previous(w http.ResponseWriter, r *http.Request) {
	if previous_fib == 0 {
		const response = "cannot go previous."
		fmt.Fprint(w, response)
	} else {
		temp := previous_fib
		previous_fib = current_fib - temp
		current_fib = temp
		var response = strconv.Itoa(previous_fib)
		fmt.Fprint(w, response)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var response = "<a href=\"/previous\">/previous</a> - For previous fibonacci number\n</br>" +
		"<a href=\"/current\">/current</a> - For current fibonacci number\n</br>" +
		"<a href=\"/next\">/next</a> - For next fibonacci number\n</br>"
	fmt.Fprint(w, response)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/current", current)
	http.HandleFunc("/previous", previous)
	http.HandleFunc("/next", next)
}

func main() {
	handleRequests()
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}
	fmt.Println("Go Web App Started on Port " + httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
