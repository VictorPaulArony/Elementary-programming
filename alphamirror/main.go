package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	data := os.Args[1]
	var res string
	for _, val := range data {
		res += string(Check(val))
	}
	for _, val := range res {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func Check(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return 'z' - (r - 'a')
	}
	if r >= 'A' && r <= 'Z' {
		return 'Z' - (r - 'A')
	}
	return r
}
