package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(eb.getGreeting())
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
