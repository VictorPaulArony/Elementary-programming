package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		return
	}

	data := os.Args[1]
	num := Atoi(data)
	// The main function checks if the number is positive and a power of 2 using the bitwise operation
	for num != 1 {
		if num%2 != 0 {
			fmt.Println("false")
			return
		}
		num /= 2

	}
	fmt.Println("true")
}

func Atoi(str string) int {
	var res int
	for _, s := range str {
		res = res*10 + int(s-'0')
	}
	return res
}
