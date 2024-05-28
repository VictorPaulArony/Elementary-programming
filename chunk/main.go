package main

import (
	"chunk/check"
	"fmt"
)

func main() {
	fmt.Println(check.Chunk([]int{}, 10))
	fmt.Println(check.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0))
	fmt.Println(check.Chunk([]int{ 1, 2, 3, 4, 5}, 2))
	fmt.Println(check.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5))
	fmt.Println(check.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4))
}
