package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		data := os.Args[len(os.Args)-1]
		for _, d := range data {
			z01.PrintRune(d)
		}
	} else {
		return
	}
	z01.PrintRune('\n')
}
