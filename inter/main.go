package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	d1 := os.Args[1]
	d2 := os.Args[2]
	res := Inter(d1, d2)
	os.Stdout.WriteString(res + "\n")
}

func Inter(str1, str2 string) string {
	c1 := make(map[rune]bool)
	c2 := make(map[rune]bool)
	var res string
	for _, s := range str2{
		if !c1[s] {
			c1[s] = true
		}
	}
	for _, s := range str1 {
		if c1[s] && !c2[s]{
			c2[s] = true
			res =res + string(s)
		}
	}
	return res
}

/*
inter
Instructions
Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

The display will be followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing.

Usage
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
$
*/
