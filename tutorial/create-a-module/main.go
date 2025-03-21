package main

import (
	"fmt"
)

func main() {
	// Call the Hello function from the greeting package.
	message := Hello("Gopher")
	fmt.Println(message)
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}