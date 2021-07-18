package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/api/", homePage)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	fmt.Println("Server Started")
	handleRequests()
}
