package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
