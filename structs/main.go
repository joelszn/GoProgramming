package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

//struct being defined
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Nelson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	// jimPointer := &jim
	jim.updatefName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updatefName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
