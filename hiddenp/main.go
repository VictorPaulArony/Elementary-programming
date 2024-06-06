package main

import (
	"os"
)

func main() {
	data1 := os.Args[1]
	data2 := os.Args[2]
	if Check(data1, data2) {
		os.Stdout.WriteString("1\n")
	} else {
		os.Stdout.WriteString("0\n")
	}
}

func Check(str1, str2 string) bool {
	count1 := 0
	count2 := 0
	l1 := len(str1)
	l2 := len(str2)
	for count1 < l1 && count2 < l2 {
		if str1[count1] == str2[count2] {
			count1++
		}
		count2++
	}
	return count1 == l1
}

/*
Instructions
Write a program named hiddenp that takes two string and that, if the first string is hidden in the second one, displays 1 followed by a newline ('\n'), otherwise it displays 0 followed by a newline.

Let s1 and s2 be string. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

If s1 is an empty string, it is considered hidden in any string.

If the number of arguments is different from 2, the program displays nothing.

Usage
$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
$ go run .
$
*/
