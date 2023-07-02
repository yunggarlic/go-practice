package main

import "log"

func main() {
	something := "Some thing"
	log.Println(something)
	changeSomethingByPointer(&something)
	log.Println(something)
}

func changeSomethingByPointer(s *string) {
	*s = "Something else"
}
