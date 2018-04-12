package main

import "fmt"

type deck []string

func (d deck) print() {
	//foreach syntax
	for i, card := range d {
		fmt.Println(i, card)
	}
}
