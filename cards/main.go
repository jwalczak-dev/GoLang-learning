package main

import "fmt"

var deckSize int

func main() {
	// type conversion - Example with creating slice of bytes from string
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	// fmt.Println("================== Deck from file ====================")
	// newDeckFromFile("my_cards").print()

	cards2 := newDeck()
	cards2.shuffle()
	cards2.print()
}

func (d deck) print2() {
	for i, card := range d {
		fmt.Println(i*2, card)
	}
}
