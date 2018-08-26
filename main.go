package main

import (
	"fmt"
)

type chatBot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(cb chatBot) {
	fmt.Println(cb.getGreeting())
}
