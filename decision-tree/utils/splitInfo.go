package utils

import (
	"math"
	"sync"
)

// function to calculate the split infomation
// measure how uniformly the data is (prevent bias)
func SplitInformation(data [][]string, columnIndex int) float64 {
	dataRows := len(data)
	if dataRows == 0 {
		return 0.0 // prevent division by 0
	}

	splitInfo := 0.0
	colType := DetermineDataType(data, columnIndex)

	if colType == "categorical" {
		splits := splitDataCategoricalParallel(data, columnIndex)
		// goroutine to calculate the entropy
		var wg sync.WaitGroup
		mu := &sync.Mutex{}

		for _, subset := range splits {
			wg.Add(1)
			go func(subset [][]string) {
				defer wg.Done()
				prob := float64(len(subset)) / float64(dataRows)
				if prob > 0 {
					mu.Lock()
					splitInfo -= prob * math.Log2(prob)
					mu.Unlock()
				}
			}(subset)
		}
		wg.Wait()
	} else {
		// Handle numeric attribute
		leftSplit, rightSplit, _ := SplitByNumeric(data, columnIndex)
		if len(leftSplit) > 0 {
			prob := float64(len(leftSplit)) / float64(dataRows)
			if prob > 0 {
				splitInfo -= prob * math.Log2(prob)
			}
		}
		if len(rightSplit) > 0 {
			prob := float64(len(rightSplit)) / float64(dataRows)
			if prob > 0 {
				splitInfo -= prob * math.Log2(prob)
			}
		}
	}

	return splitInfo
}
