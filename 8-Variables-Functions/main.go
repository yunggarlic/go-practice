package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Variables
	var whatToSay string
	whatToSay = "Goodbye, cruel world!"

	fmt.Println(whatToSay)

	var i int = 7

	fmt.Println("i --> ", i)

	whatWasSaid := saySomething() // := is a shortcut for declaring and initializing a variable, with an inferred type

	fmt.Println("whatWasSaid --> ", whatWasSaid)
}

func saySomething() string {
	return "something"
}
