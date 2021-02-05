package main

const handSize = 5

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, _ := deal(cards, handSize)
	hand.saveToFile("my_hand")
}
