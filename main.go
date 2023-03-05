package main

func main() {
	//cards := newDeck()
	vehicles := vehicle()
	first, second := multipleReturn(vehicles, 1)

	first.result()
	second.result()
}