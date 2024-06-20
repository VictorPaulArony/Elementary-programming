package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	data := os.Args[1]
	num := Atoi(data)

	if Power(num) {
		os.Stdout.WriteString("true\n")
	} else {
		os.Stdout.WriteString("false\n")
	}
}

// function Power to return true if a number is power of two
func Power(num int) bool {
	for num > 1 {
		if num%2 != 0 {
			return false
		}
		num /= 2
	}
	return true
}

// function Atoin to convert string to int
func Atoi(str string) int {
	var res int
	for _, s := range str {
		res = res*10 + int(s-'0')
	}
	return res
}

