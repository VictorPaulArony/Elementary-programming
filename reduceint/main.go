package main

import (
	"fmt"

	"reduceint/reduce"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 3}
	fmt.Println(reduce.ReduceInt(as, mul))
	fmt.Println(reduce.ReduceInt(as, sum))
	fmt.Println(reduce.ReduceInt(as, div))
}

// Expected
// 1000
// 502
// 250
