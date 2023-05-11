package main

import "fmt"

func main() {
	cards := newDeck()

	h, r := deal(cards, 5)
	fmt.Println("Hand:", h)
	fmt.Println("Remaining cards:", r)
	cards.shuffle()
	cards.print()

}
