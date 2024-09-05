package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := 0

	lines := ""
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln("INVALID FILE", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = scanner.Text()

		lines = strings.ReplaceAll(lines, "x", " ")
		nums := strings.Split(lines, " ")

		num := []int{}
		for _, val := range nums {
			digit, _ := strconv.Atoi(val)
			num = append(num, digit)
		}

		res += 2*(num[0]*num[1]) + 2*(num[1]*num[2]) + 2*(num[2]*num[0])

		sort.Ints(num)
		res += num[0] * num[1]
	}
	fmt.Printf("result: %d\n", res)
}
