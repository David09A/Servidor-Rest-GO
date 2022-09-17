package main

import (
	"net/http"
)

func main() {
	//routers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	//Start the server
	http.ListenAndServe(":3000", nil)
	println("Hello world")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world contact"))
}
