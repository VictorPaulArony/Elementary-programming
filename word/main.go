package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	result := ""
	for a := len(args)-1; a >= 0; a-- {
		if args[a] != " " {
			result = string(args[a]) + result
			
		} else {
			if result != "" &&args[a] != " "{
				result = ""
				break
			}
		}
	}
	fmt.Println(result)
}
