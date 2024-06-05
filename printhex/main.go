package main

import (
	"fmt"
	"os"
)

// the main function for data output
func main() {
	if len(os.Args) < 1 {
		return
	}
	data := os.Args[1]
	if !Vaild(data) {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(Hex(Atoi(data)))
}

// func Hex converts the int to a significant hex-decimal value
func Hex(num int) string {
	var res string
	dec := "0123456789abcdef"
	for num > 0 {
		hex := num % 16
		res = string(dec[hex]) + res
		num /= 16
	}
	return res
}

// func Atoi convert string to int
func Atoi(str string) int {
	res := 0
	for _, val := range str {
		res = res*10 + int(val-'0')
	}
	return res
}

// func Valid check if the argument is a digit 0-9
func Vaild(str string) bool {
	if str == "" {
		return false
	}
	for _, s := range str {
		if s < '0' || s > '9' {
			return false
		}
	}
	return true
}

// $ go run . 10
// a
// $ go run . 255
// ff
// $ go run . 5156454
// 4eae66
// $ go run .
// $ go run . "123 132 1" | cat -e
// ERROR$
// $
