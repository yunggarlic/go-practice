package main

import (
	"log"
	"math/rand"
)

func main() {

	random := rand.Intn(100)

	if random > 50 {
		log.Println("Wow! Over 50, you must be lucky!")
	} else {
		log.Println("Not over 50, you must be unlucky!")
	}

	switch random {
	case 0:
		log.Println("But wowie zowie, you got zero!")

	case 100:
		log.Println("Un-freaking-believable. You got 100! Are you cheating?")

	default:
		log.Println("You got a number between 1 and 99. Good job!")
	}
	log.Println("Your number was:", random)

	log.Println("--The End--")
}
