package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	data := os.Args[1]

	var res string
	for _, d := range data {
		res += string(Check(d))
		continue
	}
	for _, val := range res {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func Check(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return 'a' + (r-'a'+13)%26
	}
	if r >= 'A' && r <= 'Z' {
		return (r-'A'+13)%26 + 'A'
	}
	return r
}

// Expected
// $ go run . "abc"
// nop
// $ go run . "hello there"
// uryyb gurer
// $ go run . "HELLO, HELP"
// URYYB, URYC
// $ go run .
// $
