package main

import (
	"fmt"
)

func main() {
	var userId int
    var price float64
	var status bool
	var name string
	var (
		email       = "user@example.com"
		phoneNumber = "1111111111"
	)

	// Print type of a variable.
	fmt.Printf("%T\n", userId)
	fmt.Printf("%T\n", price)
	fmt.Printf("%T\n", status)
	fmt.Printf("%T\n", name)

	// Print value of a variable.
	fmt.Printf("%v\n", email)
	fmt.Printf("Your phone number is %[2]v. Your email id is %[1]v\n", email, phoneNumber)
}