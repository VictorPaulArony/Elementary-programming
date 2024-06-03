package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	input := os.Args[1]
	res := ""
	for _, v := range input {
		if v >= 'a' && v <= 'z' {
			for i := 'a'; i <= v; i++ {
				res += string(v)
			}
		} else if v >= 'A' && v <= 'Z' {
			for i := 'A'; i <= v; i++ {
				res += string(v)
			}
		} else {
			res += string(v)
		}
	}
	os.Stdout.WriteString(res + "\n")
}
