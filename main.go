package main

const handSize int = 5

func main() {
	cards := new_deck()
	hand, remainingDeck := deal(cards, handSize)
	hand.print()
	remainingDeck.print()
}
