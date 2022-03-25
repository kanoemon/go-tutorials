package main

import (
	"fmt"
	"log"

	"kanoemon/go-tutorials/greetings"

	"golang.org/x/example/stringutil"
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

	fmt.Println(stringutil.Reverse("Hello"))
}
