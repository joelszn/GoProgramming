package main

func main() {
	//dealing with splices
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//iterate over slice of cards

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//using function instead
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
