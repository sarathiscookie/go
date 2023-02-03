// Variables, Constants, Variables in strings, DataTypes
package main

import (
	"fmt"
	"strings"
)

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

	// --------------------- Array ------------------------------
	/*var vehicles []string
	vehicles = append(vehicles, "Mercedes Benz", "Range Rover")

	fmt.Printf("First value: %v\nArray type: %T\nArray length: %v\n", vehicles, vehicles, len(vehicles))*/

	// ------------ Loops in GO (for and foreach) ------------
	// ------------ if else continue break -----------
	var remainingSeats uint = 50
	var userNames []string

	for {
		firstName, lastName, email, ticketsBooked := getUserInputs()

		isValidName, isValidEmail, isValidBooking := validateUserInput (firstName, lastName, email, ticketsBooked)

		if isValidName == false {
			fmt.Println("First name and Last name must be minimum 2 letters.")
			break
		}

		if !isValidEmail {
			fmt.Println("Enter valid email address.")
			break
		}

		if ticketsBooked > remainingSeats {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v \n", remainingSeats, ticketsBooked)
			continue
		}

		remainingSeats = remainingSeats - ticketsBooked

		userNames = append(userNames, firstName + " " + lastName)

		if remainingSeats > 0 && isValidBooking == true {
			fmt.Printf("%v %v, Welcome to ticket booking system. Thank you for booking your tickets. No of bookings %v \n", firstName, lastName, ticketsBooked)

			fmt.Printf("Remaining tickets are %v \n", remainingSeats)
			
			fmt.Printf("You will get email confirmation on your email id %v \n", email)

			userFirstNames := getFirstNames(userNames)

			fmt.Printf("The first names of bookings are %v \n", userFirstNames)
		} else {
			fmt.Printf("Whoops! Check booking count or Seats are booked out. Come back next year. \n")
			break
		}
	}

	// --------------- Switch --------------
	/*var city string
	fmt.Printf("Enter your city: \n")
	fmt.Scan(&city)

	switch city {
	case "Newyork":
		fmt.Printf("Your native city is %v", city)
	case "London", "Berlin":
		fmt.Printf("Your native city is %v", city)
    default:
		fmt.Println("Sorry your city is not in our list.")		
	}*/
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var ticketsBooked uint

	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter no of tickets: \n")
	fmt.Scan(&ticketsBooked)

	return firstName, lastName, email, ticketsBooked
}

func validateUserInput(firstName string, lastName string, email string, ticketsBooked uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidBooking := ticketsBooked > 0

	return isValidName, isValidEmail, isValidBooking
}

func getFirstNames(userNames []string) ([]string) {
	firstNames := []string {}

	for _, userName := range userNames {
		var names = strings.Fields(userName)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}