package utils

import "math"

// function to calculate the split infomation
// measure how uniformly the data is (prevent bias)
func SplitInformation(data [][]string, targetIndex int) float64 {
	total := len(data)

	// Create a map to count occurrences of each class label
	columnCount := make(map[string]int)
	for _, row := range data {
		columnCount[row[targetIndex]]++
	}

	splitInfo := 0.0
	for _, count := range columnCount {
		prob := float64(count) / float64(total)
		splitInfo -= prob * math.Log2(prob)
	}
	return splitInfo
}
