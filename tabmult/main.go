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
	num, _ := strconv.Atoi(data)

	for i := 1; i <= 9; i++ {
		// res += i * num
		var res string
		res += strconv.Itoa(i) + " X " + strconv.Itoa(num) + " = " + strconv.Itoa(i*num) + "\n"
		os.Stdout.WriteString(res)
	}
}
