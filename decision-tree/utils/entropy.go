package utils

import (
	"math"
)

// function to salculate the entropy for the dataset
//
//	for each class lable calculate probability and entropy
func CalculateEntropy(data [][]string, targetIndex int) float64 {
	countLables := make(map[string]int)
	for _, row := range data {
		if len(row) > targetIndex { // Add bounds check
			countLables[row[targetIndex]]++
		}
	}
	var entropy float64
	dataLen := len(data)
	if dataLen == 0 {
		return 0.0 // division by 0 in prob
	}
	// Calculate the entropy based on the class label counts
	for _, count := range countLables {
		prob := float64(count) / float64(dataLen)
		// Update the entropy using the formula: -p * log2(p)
		if prob > 0 { // Avoid log2(0) which is undefined
			entropy -= prob * math.Log2(prob)
		}
	}
	// fmt.Println(entropy)
	return entropy
}

// Functin to find the index of a column by its name
func FindColumnIndex(headers []string, columnName string) int {
	for i, header := range headers {
		if header == columnName {
			return i
		}
	}
	return -1
}
