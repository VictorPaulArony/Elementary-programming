package main

import (
	"os"
)
// Ensure exactly one command-line argument is provided
func main() {
	if len(os.Args) < 2 {
		return
	}
	data := os.Args[1]
	atoi := Atoi(data)
	next := Prime(atoi)
	out := Star(next)
	os.Stdout.WriteString(out + "\n")
}
// Atoi ,converting string to int
func Atoi(str string) int {
	res := 0
	for _, char := range str {
		res = res*10 + int(char-'0')
	}
	return res
}
//Itoa converts int to string 
func Itoa(num int) string {
	res := ""
	for num > 0 {
		n := num % 10
		dig := string(rune(n + '0'))
		res = dig + res
		num /= 10
	}

	return res
}

func Star(slc []int) string {
	var res string
	for _, val := range slc {
		res = res + Itoa(val) + "*"
	}
	return res[:len(res)-1]
}
//IsPrime check if the int is a prime number 
func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for a := 2; a*a < num; a++ {
		if num%a == 0 {
			return false
		}
	}
	return true
}
// Prime get the list of all factorial prime numbers of a given number 
func Prime(num int) []int {
	var res []int
	for i := 2; i <= num; i++ {
		for IsPrime(i) && num%i == 0 {
			res = append(res, i)
			num /= i

		}
	}
	return res
}

/*Instructions
Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

Factors must be displayed in ascending order and separated by *.

If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

Usage
$ go run . 225225
3*3*5*5*7*11*13
$ go run . 8333325
3*3*5*5*7*11*13*37
$ go run . 9539
9539
$ go run . 804577
804577
$ go run . 42
2*3*7
$ go run . a
$ go run . 0
$ go run . 1
$*/
