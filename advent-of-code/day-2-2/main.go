package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res := 0

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln("INVALID FILE", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines := scanner.Text()

		lines = strings.ReplaceAll(lines, "x", " ")

		nums := strings.Split(lines, " ")

		num := []int{}
		for _, digit := range nums {
			n, _ := strconv.Atoi(digit)
			num = append(num, n)
		}

		sort.Ints(num)
		res += 2*(num[0]+num[1]) + num[0]*num[1]*num[2]

	}
	log.Printf("result: %d\n", res)
}
