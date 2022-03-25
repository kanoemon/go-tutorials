package main

import (
	"fmt"
	"log"

	"github.com/kanoemon/go-tutorials/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("kanoe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
