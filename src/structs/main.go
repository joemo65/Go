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

	fmt.Println(jim)
}
