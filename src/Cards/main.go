package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	PrintArray(cards)
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

	//foreach syntax
	for i, card := range arr {
		fmt.Println(i, card)
	}
}