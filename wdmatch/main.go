package main

import (
	"fmt"
	"os"
)

func main() {
	data1 := os.Args[1]
	data2 := os.Args[2]
	if Data(data1, data2) {
		fmt.Println(data1)
	}
}

func Data(data1, data2 string) bool {
	c1 := 0
	c2 := 0
	l1 := len(data1)
	l2 := len(data2)
	for c1 < l1 && c2 < l2 {
		if data1[c1] == data2[c2] {
			c1++
		}
		c2++
	}
	return c1 == l1
}
