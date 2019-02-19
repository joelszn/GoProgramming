package main

import "fmt"

// the word bot doesn't matter it is just a placeholder
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we can omit the value if it is not being used
func (englishBot) getGreeting() string {
	// custom logic for english greeting
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	// custom logic for english greeting
	return "Hola"
}
