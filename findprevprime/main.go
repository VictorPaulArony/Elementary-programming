package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(10))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	for i := nb; i > 1; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
