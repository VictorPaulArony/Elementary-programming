package main

import (
	"fmt"
)

func main() {
	var res string
	str := "apple banana cherry date elderberry"
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else if word != "" {
			res = word + " "+ res
			word = ""
		}
	}
	if word != " " {
		res = word + " " + res
	}
	fmt.Println(res)
}
