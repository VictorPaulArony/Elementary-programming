package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	data1 := os.Args[1]
	data2 := os.Args[2]
	os.Stdout.WriteString(Inter(data1, data2) + "\n")
}

func Inter(str1, str2 string) string {
	var res string
	count := make(map[rune]bool)
	seen := make(map[rune]bool)
	for _, val := range str2 {
		count[val] = true
	}
	for _, val := range str1 {
		if count[val] && !seen[val] {
			seen[val] = true
			res = res+ string(val)
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