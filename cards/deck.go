package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	/*
	* Basically the approach here is to create two different slices
	* One of suits
	* Another of possible numbers
	* then loops the values into a deck instead of creating manually
	* an entire deck :D
	 */

	// the reason these are not of deck type because they're just temp
	// strings to create the deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four"}

	// , "Five","Six", "Seven"
	// ,"Seven","Eight","Nine", "Ten"
	// }

	// instead of using i & j in Go we can use '_'
	for _, suit := range cardSuits {
		for _, num := range cardNumbers {
			cards = append(cards, num+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// TODO
func shuffle() {

}

//we can return two separate things from a fcn in Go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// s the slice of string from the split
	// s is by convention
	s := strings.Split(string(bs), ",") // Ace of Spades,Three of Clubs,
	return deck(s)
}

// helper fcn will take a deck and return string representation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}
