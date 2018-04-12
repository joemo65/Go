package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jhalpert@dundermifflin.com",
			zipCode: "94000",
		},
	}

	jim.print()

	jimPointer := &jim
	jimPointer.updateName("Michael", "Scott")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
