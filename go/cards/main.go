package main

func main() {
	cards := newDeck()

	/* hand, remainingCards := deal(cards, 5)

	hand.printDeck()
	remainingCards.printDeck() */

	cards.saveToFile("myDeck")
}
