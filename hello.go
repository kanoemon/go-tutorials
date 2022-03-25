package main

import (
	"fmt"

	"github.com/kanoemon/go-tutorials/greetings"
)

func main() {
	message := greetings.Hello("kanoe")
	fmt.Println(message)
}
