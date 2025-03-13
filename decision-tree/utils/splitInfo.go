package utils

import (
	"math"
	"sync"
)

// function to calculate the split infomation
// measure how uniformly the data is (prevent bias)
func SplitInformation(data [][]string, targetIndex int) float64 {
	dataRows := len(data)
	if dataRows == 0 {
		return 0.0
	}

	splitInfo := 0.0

	colType := DetermineDataType(data, targetIndex)
	if colType == "categorical" {
		splits := SplitDataCategorical(data, targetIndex)
		// goroutine to calculate the entropy
		var wg sync.WaitGroup
		mu := &sync.Mutex{}

		for _, subset := range splits {
			wg.Add(1)
			go func(subset [][]string) {
				defer wg.Done()

				prob := float64(len(subset)) / float64(dataRows)
				splitInfo -= prob * math.Log2(prob)

				mu.Unlock()
			}(subset)
		}

	}else {
		// Handle numeric attribute
		leftSplit, rightSplit, _ := splitByNumeric(data, targetIndex)
		
		if len(leftSplit) > 0 {
			prob := float64(len(leftSplit)) / float64(dataRows)
			splitInfo -= prob * math.Log2(prob)
		}
		if len(rightSplit) > 0 {
			prob := float64(len(rightSplit)) / float64(dataRows)
			splitInfo -= prob * math.Log2(prob)
		}
	}

	return splitInfo
}
