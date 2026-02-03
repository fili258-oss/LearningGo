package main

import (
	"fmt"
	"example.com/phrases"
	"log"
)

func main() {
	fmt.Println("HI")
	log.SetPrefix("Phrases: ")
	log.SetFlags(0)

	//days with key number day 1,2,3,4,5,6,7
	keys := []int{1,3,4,6,4,7,5}

	messages, err := phrases.Days(keys)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}