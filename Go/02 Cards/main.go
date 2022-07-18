package main

func main() {
	cards := newDeckFromFile("remainingDeck.json")
	cards.shuffle()

	cards.print()
}

func createHand(d deck) (deck, deck) {
	return deal(d, 2)
}

func saveDeck(hand deck, reamining deck) {
	hand.saveToFile("hand")
	reamining.saveToFile("remainingDeck")
}
