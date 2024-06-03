package main

import "os"

func main() {
	if len(os.Args) < 2 {
		return
	}
	data := os.Args[1]
	num := Atoi(data)
	sum := Sum(num)
	str := Itoa(sum)
	os.Stdout.WriteString(str + "\n")
}

func Prime(num int) bool {
	if num <= 1 {
		return false
	}
	for n := 2; n*n <= num; n++ {
		if num%n == 0 {
			return false
		}
	}
	return true
}

func Atoi(str string) int {
	res := 0
	for _, s := range str {
		res = res*10 + int(s-'0')
	}
	return res
}

func Itoa(num int) string {
	res := ""
	for num > 0 {
		n := num % 10
		dig := string(n + '0')
		res = dig + res
		num /= 10
	}
	return res
}

func Sum(num int) int {
	res := 0

	for i := 2; i <= num; i++ {
		if Prime(i) {
			res += i
		}
	}
	return res
}
