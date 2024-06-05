package main

import "github.com/01-edu/z01"

func main() {
	for a := 'a'; a <= 'z'; a++ {
		if a%2 != 0 {
			z01.PrintRune(a - ('a' - 'A'))
		} else {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
	for a := 'z'; a >= 'a'; a-- {
		if a%2 == 0 {
			z01.PrintRune(a - ('a' - 'A'))
		} else {
			z01.PrintRune(a)
		}
	}

	z01.PrintRune('\n')
}

// Instructions
// Write a program that:

// first prints the Latin alphabet alternatively in uppercase and lowercase in order (from 'A' to 'z') on a single line.
// second prints the Latin alphabet alternatively in uppercase and lowercase in reverse order (from 'Z' to 'a') on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').

// Usage
// $ go run .
// AbCdEfGhIjKlMnOpQrStUvWxYz
// ZyXwVuTsRqPoNmLkJiHgFeDcBa
// $
