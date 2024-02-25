package main

var deckSize int

func main() {
	// card := newCard()
	// fmt.Println(card)
	// printState()

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
