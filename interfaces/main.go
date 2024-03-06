package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

// type polishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	// pb := polishBot{}

	printGreeting(eb)
	printGreeting(sb)
	// printGreeting(pb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
