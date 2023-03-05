package main

import (
	"path"
	"fmt"
)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}