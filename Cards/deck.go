package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// cardsString := ""

	// for i, card;= range d{
	// 	cardsString += fmt.Sprintf("%v ", i+1) + card + "\n"
	// }

	// return cardsString

	return strings.Join([]string(d), ",")
}

// saveToFile writes the deck to a file with the given filename.//+
// Parameters://+
//   - filename: A string representing the name of the file to save the deck to.//+
//
// Returns://+
//   - An error if the write operation fails, otherwise nil.//+
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}
