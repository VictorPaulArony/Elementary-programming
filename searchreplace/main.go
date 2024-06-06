package main

import (
	"os"
)

func main() {
	data := os.Args[1]
	s := os.Args[2]
	r := os.Args[3]
	res := ""
	for _, val := range data {
		if string(val) == s {
			res += r
		} else {
			res += string(val)
		}
	}
	os.Stdout.WriteString(res + "\n")
}

// Instructions
// Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

// If the number of arguments is different from 3, the program displays nothing.

// If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

// Usage
// $ go run . "hella there" "a" "o"
// hello there
// $ go run . "hallo thara" "a" "e"
// hello there
// $ go run . "abcd" "z" "l"
// abcd
// $ go run . "something" "a" "o" "b" "c"
// $
