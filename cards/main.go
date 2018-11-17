package main

import (
	"fmt"
)

func main() {
	//dealing with splices
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//iterate over slice of cards

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
