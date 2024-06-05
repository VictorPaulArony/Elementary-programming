package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	data := os.Args[1]
	atoi := Atoi(data)
	os.Stdout.WriteString(Itoa(atoi) + "\n")
}

func Atoi(str string) int {
	res := 0
	for _, num := range str {
		res = res*10 + int(num-'0')
	}
	return res
}

func Itoa(num int) string {
	var res string
	for num > 0 {
		n := num % 10
		res = string(n+'0') + res
		num /= 10

	}
	return res
}

// Instructions
// Write a program that takes a number as argument, and prints it in binary value without a newline at the end.

// If the the argument is not a number, the program should print 00000000.
// Usage :
// $ go run . 1
// 00000001$
// $ go run . 192
// 11000000$
// $ go run . a
// 00000000$
// $ go run . 1 1
// $ go run .
// $
