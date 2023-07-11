package main

import (
	"fmt"
	/*"math"
	"time"*/
)

/*func sum(a int, b int) int {
	return a + b
}*/

/*func multiply(a, b int) int {
	return a * b
}*/

func swap(a, b string) (string, string) {
	return a, b
}

func main() {
	/*fmt.Println("Math Round", math.Round(10.5))
	fmt.Println("Month", time.Now().Month())
	fmt.Println("Math Sqrt", math.Sqrt(2*2))
	fmt.Println("Math Pi", math.Pi)*/
	/*fmt.Println("Sum:", sum(4, 5))*/
	/*fmt.Println("Multiply:", multiply(4, 5))*/
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
