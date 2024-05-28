package main

import (
	"os"

	"github.com/01-edu/z01"
)

//"os"
//"github.com/01-edu/z01"

func main() {
	if len(os.Args) > 2 {
		os.Exit(0)
	}
	data := os.Args[1]
	res := ""
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != ' ' {
			res = string(data[i]) + res
		} else if data[i] == ' ' && res != "" {
			break
		}
	}
	if res != "" {
		for _, char := range res {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
