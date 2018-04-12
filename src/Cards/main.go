package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("hand")
	hand.print()

	fmt.Println()

	fmt.Println("remaining Deck")
	remainingDeck.print()
}
