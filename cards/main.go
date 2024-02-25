package main

import "fmt"

var deckSize int

func main() {
	card := newCard()
	fmt.Println(card)
	printState()
}

func newCard() string {
	return "Five of Diamonds"
}
