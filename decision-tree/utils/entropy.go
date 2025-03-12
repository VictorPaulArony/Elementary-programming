package utils

import "math"


// function to salculate the entropy for the dataset
//  for each class lable calculate probability and entropy
func CalculateEntropy(data [][]string, targetIndex int) float64 {
	var entropy float64
	dataLen := len(data)
	if dataLen == 0 {
		return 0
	}
	
	// Create a map to count occurrences of each class label
	countLables := make(map[string]int)
	for _, row := range data {
		countLables[row[targetIndex]]++
	}

	// Calculate the entropy based on the class label counts
	for _, count := range countLables {
		prob := float64(count) / float64(dataLen)
		// Update the entropy using the formula: -p * log2(p)
		entropy -= prob * math.Log2(prob)
	}
	return entropy
}
