package main

import (
	"fmt"
)

//create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() {

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// TODO
func shuffle() {

}

// TODO
func deal() {

}

// TODO
func saveToFile() {

}

// TODO
func newDeckFromFile() {

}
