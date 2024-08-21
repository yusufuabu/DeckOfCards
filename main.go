package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
