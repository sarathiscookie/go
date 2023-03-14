package main

import (
	"fmt"
)

func main() {
	const (
		spring = iota + 3 //3
		_ //4
		_ //5
		summer = iota + 3 //6
		_ //7
		_ //8
		fall = iota + 3 //9
		_ //10
		_ //11
		winter = iota + 3 //12
	)
	fmt.Println(winter, spring, summer, fall)
}