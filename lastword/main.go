package main

import (
	"fmt"
	//"os"
	//"github.com/01-edu/z01"
)

func main() {
	str := "    FOR GOPHERS res     "
	// last := len(str)-1
	// fmt.Println(last)
	var s []string
	sj := ""
	for _, char := range str {
		if char == ' ' && len(sj) != 0 {
			s = append(s, sj)
			sj = ""
			continue
		} else if char != ' ' {
			sj += string(char)
		}
	}
	if len(sj) != 0 {
		s = append(s, sj)
	}
	fmt.Println(s[len(s)-1])
	println(len(s))
}
