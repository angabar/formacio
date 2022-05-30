package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (singleDeck deck) printDeck() {
	for index, card := range singleDeck {
		fmt.Println(index, card)
	}
}

func deal(singleDeck deck, handSize int) (deck, deck) {
	return singleDeck[:handSize], singleDeck[handSize:]
}

func (singleDeck deck) toString() string {
	return strings.Join([]string(singleDeck), ",")
}

func (singleDeck deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(singleDeck.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, fileError := ioutil.ReadFile(fileName)

	if fileError != nil {
		fmt.Println("Error: ", fileError)
		os.Exit(1)
	}

	splittedStringDeck := strings.Split(string(byteSlice), ",")

	return deck(splittedStringDeck)
}

func (singleDeck deck) shuffle() {
	for index := range singleDeck {
		newPosition := rand.Intn(len(singleDeck) - 1)
		singleDeck[index], singleDeck[newPosition] = singleDeck[newPosition], singleDeck[index]
	}
}
