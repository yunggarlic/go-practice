package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page")

}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page")

}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
