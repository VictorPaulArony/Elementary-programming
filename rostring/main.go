package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	data := os.Args[1]
	fmt.Println(Convert(data))
}

func split(str string) []string {
	var res []string
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else if word != "" {
			res = append(res, word)
			word = ""
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}

func Convert(str string) string {
	s := split(str)
	if len(s) == 0 {
		return ""
	}
	res := append(s[1:], s[0])
	return join(res)
}

func join(str []string) string {
	var res string
	for i, s := range str {
		if i > 0 {
			res += " "
		}
		res += s
	}
	return res
}

// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
// $
// $
