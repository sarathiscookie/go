package main

import "fmt"

func main() {
	/***************  Variable & Data Type *************/
	/*fmt.Println("Hello")

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

	// Array Type
	var ar = [4]int{120, 300, -45, 250}
	fmt.Printf("%T \n", ar)

	// Slice Type
	var sl = []string{"Mumbai", "Delhi", "GOA"}
	fmt.Printf("%T \n", sl)
	
	// Map Type
	var ma = map[string]float64{
		"price": 12345.1,
		"total": 67890.2,
	}
	fmt.Printf("%T \n", ma)

	// Struct Type
	type Person struct {
		name string
		age int
	}

	var you Person

	fmt.Printf("%T \n", you)

	// Pointer Type
	var p int = 2
	pp := &p
	fmt.Printf("%T \n", pp)

	// Interface Type & Channel Type
	// TODO: Later 

	// Function Type
	fmt.Printf("%T \n", base)*/

    /*********************** Operators *********************/
	// Arithmetic
	ar1, ar2 := 4, 2
	tar := (ar1 + ar2) * (ar1 - ar2) / 2
	fmt.Println(tar)

	// Assignment
	as1, as2 := 5, 6
	as1 += as2
	as1 *= 10
	as1 -= 2
	as1 /= 2
	as1 %= 2
	fmt.Println(as1)

	// Comparison
	cm1, cm2 := 10, 10
	fmt.Println(cm1 == cm2)
	fmt.Println(cm1 != cm2)
	fmt.Println(cm1 == 10 && cm2 < 11)
}

func base() {
	//
}
