package main

import (
	"fmt"
	"github.com/thenakulchawla/expert-octo-funicular/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Nakul")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
