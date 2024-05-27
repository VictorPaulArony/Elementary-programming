package main

import (
	"fmt"

	"strlen/piscine"
)

func main() {
	l := piscine.StrLen("hėllo world")
	// l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
