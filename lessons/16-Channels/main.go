package main

import (
	"log"
	"math/rand"
)

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

const numPool = 55

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan) // close the channel when we are done

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
