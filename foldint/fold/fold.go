package fold

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, v := range a {
		n = f(n, v)
	}
	fmt.Println(n)
}
