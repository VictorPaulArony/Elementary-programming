package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	data := os.Args[1]
	vol := "aeiouAEIOU"
	word := ""
	res := ""
	check := false
	for i, s := range data {
		for _, v := range vol {
			if s == v {
				check = true
				res += data[i:] + word + "ay"
			}
			
		}
		word += string(s)
	}
	if !check{
		os.Stdout.WriteString("No vowels\n")
		return
	}
	os.Stdout.WriteString(res + "\n")
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
