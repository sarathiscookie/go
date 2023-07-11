package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Math Round", math.Round(10.5))
	fmt.Println("Month", time.Now().Month())
	fmt.Println("Math Sqrt", math.Sqrt(2*2))
	fmt.Println("Math Pi", math.Pi)
}
