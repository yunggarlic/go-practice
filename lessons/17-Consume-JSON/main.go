package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func main() {
	myJson := `[
	{
		"firstName": "John",
		"lastName": "Doe",
		"age": 25
	}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error:", err)
	}
	log.Printf("Unmarshalled: %v", unmarshalled)

	var people []Person

	var m1 Person
	m1.FirstName = "John"
	m1.LastName = "Doe"
	m1.Age = 25

	people = append(people, m1)

	var m2 Person
	m2.FirstName = "Mary"
	m2.LastName = "Doe"
	m2.Age = 30

	people = append(people, m2)

	log.Printf("Marshalled: %v", people)

	newJson, err := json.MarshalIndent(people, "", "    ")
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println(string(newJson))
}
