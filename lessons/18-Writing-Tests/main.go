package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Quotient:", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	log.Println("Dividing", x, "by", y)

	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}

	result = x / y
	return result, nil
}
