package main

import (
	"fmt"
	"log"

	"github.com/kanoemon/go-tutorials/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"kanoe", "taro", "hanako"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
