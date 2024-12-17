package main

import "fmt"

// Create a new type of 'deck
// which is a slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}

	cardsSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsvalues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardsSuits {
		for _, value := range cardsvalues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
