package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		return
	}

	data := os.Args[1]
	num := Atoi(data)
	// The main function checks if the number is positive and a power of 2 using the bitwise operation
	for num > 2 {
		if num%2 != 0 {
			os.Stdout.WriteString("false" + "\n")
			return
		}
		num /= 2

	}
	
	os.Stdout.WriteString("true" + "\n")
}

func Atoi(str string) int {
	res := 0
	for _, s := range str {
		res = res*10 + int(s-'0')
	}
	return res
}
