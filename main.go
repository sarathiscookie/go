package main

import "fmt"

func main() {
	cards := [] string {"Five of Diamonds \n", sample()}

	cards = append(cards, "Six of Spades \n")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func sample() string {
	return "Ace of Diamonds"
}