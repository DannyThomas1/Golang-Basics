package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// putting a '_' lets you ignore variables that are required, in this case we ignore the index value
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// A receiver function, any variable of type "deck" now gets access to the "print" method
// variable "d" is the actual copy of the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (deck,deck) tells GO that we're returning two values with a type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //from beginning to handsize, handsize to the end of deck
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//Slice of strings
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	//To create a random generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
