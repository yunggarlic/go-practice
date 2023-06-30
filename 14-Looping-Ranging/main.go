package main

import "log"

func main() {

	// classic for loop
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"cat", "human", "fish"}

	for _, animal := range animals { // blank identifier for unnecessary index
		log.Println(animal)
	}

	cats := make(map[string]string)
	cats["cat"] = "Mister Kitty"
	cats["lion"] = "Simba"
	cats["tiger"] = "Shere Khan"

	for key, cat := range cats {
		log.Println(key, ":", cat)
	}

	// var firstLine = "It was the best of times"
	// firstLine = "It was the worst of times"

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }
}
