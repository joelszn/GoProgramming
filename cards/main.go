package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_")
	cards.print()

}

//no longer needed below

// func newCard() string {
// 	return "Five of Diamonds"
// }
