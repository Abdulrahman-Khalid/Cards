package main

import "fmt"

type deck []string

//receivers

func new_deck() deck {
	newDeck := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{
		"Ace", "Two", "Three",
		"Four", "Five", "Six",
		"Seven", "Eight", "Nine",
		"Ten", "Queen", "King", "Jack"}
	for _, value := range values {
		for _, suit := range suits {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}
	return newDeck
}

func (cards deck) print() {
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}
