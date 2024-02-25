package main

import "fmt"

var deckSize int

func main() {
	// card := newCard()
	// fmt.Println(card)
	// printState()

	cards := newDeck()

	// fmt.Println(cards)

	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	// cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	cards.print()
}

func (d deck) print2() {
	for i, card := range d {
		fmt.Println(i*2, card)
	}
}
