package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	data := os.Args[1]
	num, _  := strconv.Atoi(data)
	
	for i := 1; i <= 9; i++ {
		os.Stdout.WriteString(strconv.Itoa(i) + " x " + strconv.Itoa(num) + " = " + strconv.Itoa(num*i) + "\n")
	}
}
