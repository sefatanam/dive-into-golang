package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("Sefat Anam")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Sefat", "Anan"}
	messages, err := greetings.Hellos((names))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
