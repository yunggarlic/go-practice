package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "555-555-5555",
		Age:         25,
		BirthDate:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	log.Println(user)
}
