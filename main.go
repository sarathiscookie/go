// Variables, Constants, Variables in strings, DataTypes
package main

import "fmt"

func main() {
	// Eg: Variables & Constants
	/*var conferenceName string = "Go Conference"
	
	const conferenceTickets int = 50

	var remainingTickets uint = 50 

	// Eg: Variables in strings
	fmt.Println("Welcome to",conferenceName,"booking application!")
	fmt.Printf("We have total of %v tickets and %v are still available \n\n", conferenceTickets, remainingTickets)*/

	// ----------------- DataTypes -------------------------
	/*var companyName string

	var totalEmployees int

	companyName = "TEST COMPANY"

	totalEmployees = 140000

	fmt.Printf("Company Name %v total employees %v.\n\n", companyName, totalEmployees)*/

	// -------------------- Pointer -------------------------
	/*var firstName string
	var lastName string
	var email string
	var userTickets uint
	var remainTickets uint = 50

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

    fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	remainTickets = remainTickets - userTickets 

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n\n", firstName, lastName, userTickets, email)

	fmt.Printf("Remaining tickets are %v \n", remainTickets)*/

	// Array
	var vehicles []string
	vehicles = append(vehicles, "Mercedes Benz", "Range Rover")

	fmt.Printf("First value: %v\nArray type: %T\nArray length: %v\n", vehicles, vehicles, len(vehicles))
}