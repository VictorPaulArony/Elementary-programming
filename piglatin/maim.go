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
	fmt.Println(Check(data))
}

func Vowel(runes rune) bool {
	vol := "aeiouAEIOU"
	for _, v := range vol {
		if runes == v {
			return true
		}
	}
	return false
}

func Check(data string) string {
	// Find the first vowel in the word
	var res string
	var prefix string
	for i, d := range data {
		if Vowel(d) {
			res = data[i:] + prefix + "ay"
			//break
		}
		prefix += string(d)
	}
	return res
}

// $ go run .
// $ go run . pig | cat -e
// igpay$
// $ go run . Is | cat -e
// Isay$
// $ go run . crunch | cat -e
// unchcray$
// $ go run . crnch | cat -e
// No vowels$
// $ go run . something else | cat -e
// $
