package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [] string {
		"Clubs", "Diamonds", "Spades", "Hearts",
	}

	cardValues := [] string {
		"Ace", "1", "2", "3",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit +" of "+ value)
		}
	}

	return cards
}

func (d deck) list() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}