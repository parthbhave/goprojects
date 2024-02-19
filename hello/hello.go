package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Bhave")
	fmt.Println(message)
}
