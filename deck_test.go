package main

import (
	"os"
	"testing"
)

const numOfCards = 52
const firstCard = "Ace of Spades"
const lastCard = "Jack of Clubs"

func TestCards(t *testing.T) {
	cards := newDeck()
	if len(cards) != numOfCards {
		t.Error("Expected deck's length of", numOfCards, "but got", len(cards))
	}

	if cards[0] != firstCard {
		t.Error("Expected first card to be"+firstCard+" but got", cards[0])
	}

	if cards[len(cards)-1] != lastCard {
		t.Error("Expected last card to be "+lastCard+" but got", cards[len(cards)-1])
	}
}

func TestIODeck(t *testing.T) {
	const fileName = "_deckTesting"
	os.Remove(fileName)
	newDeck := newDeck()
	newDeck.saveToFile(fileName)
	loadedDeck := loadFromFile(fileName)
	if len(loadedDeck) != numOfCards {
		t.Error("Expected deck's length of", numOfCards, "but got", len(loadedDeck))
	}
	os.Remove(fileName)
}
