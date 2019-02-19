package main

import (
	"fmt"
)

func main() {
	// slice of 10 ints print if index is even or odd

	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range numSlice {
		if numSlice[i]%2 == 0 {
			fmt.Println(numSlice[i], "is even")
		} else {
			fmt.Println(numSlice[i], "is odd")
		}
	}
}
