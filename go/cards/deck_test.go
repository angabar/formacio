package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(deck))
	}

	if deck[0] != "Spades of Ace" {
		t.Errorf("Expected first card Spades of Ace but got: %v", deck[0])
	}

	if deck[len(deck)-1] != "Clubs of Four" {
		t.Errorf("Expected first card Clubs of Four but got: %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"

	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got: %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
