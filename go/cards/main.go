package main

func main() {
	cards := newDeck()

	cards.shuffle()

	/* hand, remainingCards := deal(cards, 5)

	hand.printDeck()
	remainingCards.printDeck() */

	cards.saveToFile("myDeck")

	/* deckFromFile := newDeckFromFile("myDeck")
	deckFromFile.printDeck() */
}
