package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

//PrintArray prints each element of a string array
func PrintArray(arr []string) {
	//simple syntax
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
