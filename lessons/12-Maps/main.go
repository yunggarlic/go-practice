package main

import "log"

type User struct {
	FirstName string
}

func main() {
	myMap := make(map[string]string) //key is string, value is string
	myMap["cat"] = "Mister Kitty"
	myMap["dog"] = "Fido"
	myMap["dog"] = "Rover" //overwrites the previous value

	myStringIntMap := make(map[string]int) //key is string, value is int
	myStringIntMap["one"] = 1
	myStringIntMap["two"] = 2

	myMeMap := make(map[string]User) //key is string, value is User
	myMeMap["me"] = User{
		FirstName: "Tim",
	}

	// Slices

	var mySlice []string = []string{"one", "two", "three"}
	mySlice = append(mySlice, "one")
	log.Println(mySlice)
}
