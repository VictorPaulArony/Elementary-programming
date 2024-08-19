package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	data := os.Args[1]
	s := os.Args[2]
	r := os.Args[3]
	res := ""
	// Replace the first occurrence of s with r in data
	for _, ru := range data {
		if string(ru) == s {
			res += r
		} else {
			res += string(ru)
		}
	}

	os.Stdout.WriteString(res + "\n")
}
