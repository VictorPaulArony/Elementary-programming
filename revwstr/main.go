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
	fmt.Println(fin(data))
}

func fin(str string) string {
	s := split(str)
	Con(s)
	return join(s)
}

func split(str string) []string {
	var res []string
	var word string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else if word != "" {
			res = append(res, word)
			word = ""
		}
	}
	if word != ""{
		res = append(res, word)
	}
	return res
}

func join(res1 []string) string {
	var res string
	for i, s := range res1 {
		if i > 0 {
			res += " "
		}
		res += s
	}

	return res
}

func Con(res1 []string) {
	for i := 0; i < len(res1)/2; i++ {
		j := len(res1) - i - 1
		res1[i], res1[j] = res1[j], res1[i]
	}
}

// $ go run . "the time of contempt precedes that of indifference"
// indifference of that precedes contempt of time the
// $ go run . "abcdefghijklm"
// abcdefghijklm
// $ go run . "he stared at the mountain"
// mountain the at stared he
// $ go run . "" | cat-e
// $
// $
