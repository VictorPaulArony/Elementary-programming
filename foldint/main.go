package main

import (
	"fmt"

	"foldint/fold"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	fold.FoldInt(Add, table, ac)
	fold.FoldInt(Mul, table, ac)
	fold.FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	fold.FoldInt(Add, table, ac)
	fold.FoldInt(Mul, table, ac)
	fold.FoldInt(Sub, table, ac)
}

func Add(n1, n2 int) int {
	return n1 + n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Sub(n1, n2 int) int {
	return n1 - n2
}
