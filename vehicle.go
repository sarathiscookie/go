package main

import "fmt"

type data []string

func vehicle() data {

	vehicles := data{}

	models := [] string {
		"BMW", "Audi", "Benz",
	}

	engines := [] string {
		"Petrol", "Diesel",
	}

	for _, model := range models {
		for _, engine := range engines {
			vehicles = append(vehicles, model+"-"+engine)
		}
	} 

	return vehicles
}

func (d data) result() {
	for i, res := range d {
		fmt.Println(i, res)
	}
}

func multipleReturn(d data, size int) (data, data) {
	return d[:size], d[size:]
}