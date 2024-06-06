package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	m := len(os.Args[1:])
	res := ""
	for m > 0 {
		x := m % 10
		dig := string(rune(x + '0'))
		res = dig + res
		m /= 10
	}
	for _, val := range res {
		z01.PrintRune(val)
	}
	z01.PrintRune(10)
}
