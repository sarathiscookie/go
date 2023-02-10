package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, ticketsBooked uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidBooking := ticketsBooked > 0

	return isValidName, isValidEmail, isValidBooking
}