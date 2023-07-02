package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page")

}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(12, 17)
	fmt.Fprintf(w, fmt.Sprintf("I am %v years old", sum))

}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	var quotient float32
	if y == 0 {
		return quotient, errors.New("Cannot divide by zero")
	}
	quotient = x / y
	return quotient, nil
}

// Adds two ints and returns the sum
func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
