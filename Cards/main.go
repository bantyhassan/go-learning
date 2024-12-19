package main

import "fmt"

func main() {
	fmt.Println("Build Cards Game!!")
	// cards := newCard()
	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}

	cards := newDeck()

	// cards := newDeckFromFile("myCards")

	cards.shuffle()

	cards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("myCards")
	// cards.newDeckFromFile("myCards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	// fmt.Println(hand)
	// fmt.Println(remainingCards)

	// // for i, card := range cards {
	// // 	fmt.Println(i, card)
	// // }

	// cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
