package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) shuffle() {
	//addressing source randomization issue
	source := rand.NewSource(time.Now().UnixNano()) // time.Now().UnixNano() generates a different int64 # everytime we start up our program

	// assigning var r to new random source
	// in Go when using random it generates it from source and if the source is the same the numbers will be also
	r := rand.New(source)

	// looping thru deck slice
	for i := range d {
		// random # being generated
		// len(d) returns length of slice
		newPosition := r.Intn(len(d) - 1)

		// take whatever is at newPosition and assign it to i
		// take whatever is at index and assign it to newPosition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//we can return two separate things from a fcn in Go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	// receiving d of type deck and we're asking when the fcn is called for a filename parameter
	//then we return an error type but doesn't have to be one

	//we create the filename, and conver that into a slice of type byte code
	//0666 is just os permission code that gives all permissions

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	// byteslice is set to the filename input
	// err stores error message
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
