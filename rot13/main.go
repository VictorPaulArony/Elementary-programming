package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	data := os.Args[1]
	res := ""
	for _, v := range data {
		res += Check(v)
	}
	os.Stdout.WriteString(res + "\n")
}

func Check(r rune) string {
	res := ""
	if r >= 'a' && r <= 'z' {
		res = string((r-'a'+13)%26 + 'a')
	} else if r >= 'A' && r <= 'Z' {
		res = string((r-'A'+13)%26 + 'A')
	} else {
		res = string(r)
	}
	return res
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
