package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 || len(os.Args) > 3 {
		return
	}
	data1 := os.Args[1]
	data2 := os.Args[2]
	fmt.Println(Gcd(Atoi(data1), Atoi(data2)))
}

func Gcd(n1 int, n2 int) int {
	if n2 != 0 {
		return Gcd(n2, n1%n2)
	} else {
		return n1
	}
}

func Atoi(str string) int {
	var res int
	for _, s := range str {
		res = res*10 + int(s-'0')
	}
	return res
}

// $ go run . 42 10 | cat -e
// 2$
// $ go run . 42 12
// 6
// $ go run . 14 77
// 7
// $ go run . 17 3
// 1
// $ go run .
// $ go run . 50 12 4
// $
