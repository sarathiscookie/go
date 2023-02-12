package main

import "fmt"

type deck []string

func (d deck) list() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}