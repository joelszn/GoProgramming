package main

import (
	"fmt"
)

//struct being defined
type person struct {
	firstName string
	lastName  string
}

func main() {
	joel := person{"Joel", "Duran"}

	//OR
	//better syntax for defining a struct
	jason := person{firstName: "Jason", lastName: "Mraz"}

	//OR

	var peter person

	//updating the values on the struct

	peter.firstName = "Peter"
	peter.lastName = "Perez"

	fmt.Println(joel, jason)
	//will print out all the different fieldnames & values
	fmt.Printf("%+v", peter)

}
