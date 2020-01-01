package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

//receivers

func newDeck() deck {
	newCards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{
		"Ace", "Two", "Three",
		"Four", "Five", "Six",
		"Seven", "Eight", "Nine",
		"Ten", "Queen", "King", "Jack"}
	for _, value := range values {
		for _, suit := range suits {
			newCards = append(newCards, value+" of "+suit)
		}
	}
	return newCards
}

func (cards deck) print() {
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (cards deck) DeckToString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(cards.DeckToString()), 0666)
}

func loadFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (cards deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(src)

	for i := range cards {
		newPos := newRand.Intn(len(cards) - 1)
		cards[i], cards[newPos] = cards[newPos], cards[i]
	}
}
