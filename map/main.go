package main

import "fmt"

func main() {
	// declaring a map that all the keys are of type string & same for values
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
