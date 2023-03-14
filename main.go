package main

import (
	"fmt"
)

func main() {
	const (
		monday = iota + 1
		tuesday = iota + 1
		wednesday = iota + 1
		thursday = iota + 1
		friday = iota + 1
		saturday = iota + 1
		sunday = iota + 1
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}