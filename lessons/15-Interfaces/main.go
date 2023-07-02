package main

import "log"

type Animal interface {
	Cuteness() string
}

type Cat struct {
	Name   string
	Age    int
	IsEvil bool
}

type Giraffe struct {
	Name             string
	Age              int
	NeckLengthInches int
}

func main() {
	cat := Cat{
		Name:   "Mister Kitty",
		Age:    3,
		IsEvil: true,
	}

	giraffe := Giraffe{
		Name:             "Geoffrey",
		Age:              10,
		NeckLengthInches: 60,
	}

	log.Println("Cat cuteness:", cat.Cuteness())
	log.Println("Giraffe cuteness:", giraffe.Cuteness())

}

func (c *Cat) Cuteness() string {
	return "Super cute"
}

func (g *Giraffe) Cuteness() string {
	return "Uniquely cute"
}
