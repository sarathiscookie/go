package main

import "fmt"

func main() {
	fmt.Println("Hello")

	// Example 1
	var a int = 4

	var b float64 = 5.2

	a = int(b)

	fmt.Println(a, b)

	// Example 2
	var (
		age int
		weight float64
		status bool
	)

	fmt.Println(age, weight, status)

	// Print out the result with 3 decimal point.
	pa := 2.3
	pb := 4.5

	fmt.Printf("%.3f \n", pa * pb)

	// Constants
	const n, m int = 10, 5
	fmt.Println(n / m)

	const (
		min1 int = 250
		max1 int = 500
	)
	fmt.Println(min1, max1)

	// iota (Constant generator)
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		cc1 = -(iota + 2) // -2
		_                 // -3 (Skipped)
		cc2               // -4
		cc3               // -5
	)	
	fmt.Println(cc1, cc2, cc3)

	// Typed and Untyped constant
	const val1 = 5
	const val2 float64 = 6.25
	fmt.Println(val1 * val2)

	// Array
	var ar = [4]int{120, 300, -45, 250}
	fmt.Printf("%T \n", ar)

	// Slice
	var sl = []string{"Mumbai", "Delhi", "GOA"}
	fmt.Printf("%T \n", sl)
	
	// Map
	var ma = map[string]float64{
		"price": 12345.1,
		"total": 67890.2,
	}
	fmt.Printf("%T \n", ma)
}
