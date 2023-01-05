// Variables, Constants, Variables in strings, DataTypes
package main

import "fmt"

func main() {
	// Eg: Variables & Constants
	var conferenceName string = "Go Conference"
	
	const conferenceTickets int = 50

	var remainingTickets uint = 50 

	// Eg: Variables in strings
	fmt.Println("Welcome to",conferenceName,"booking application!")
	fmt.Printf("We have total of %v tickets and %v are still available \n\n", conferenceTickets, remainingTickets)
}