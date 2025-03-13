package utils

import (
	"sync"
)

// function to calculate the information gain
// information Gain = Entropy(parent) - (Weighted Average) * Entropy(child)
func CalculateInfoGain(data [][]string, attrName string, targetName string, headers []string) float64 {
	entropyBeforeSplit := CalculateEntropy(data, targetName, headers)
	targetIndex := findColumnIndex(headers, targetName)
	dataRows := len(data)

	if dataRows == 0 {
		return 0.0
	}

	colType := DetermineDataType(data, headers)

	entropyAfterSplit := 0.0

	if colType == "categorical" {
		splits := SplitDataCategorical(data, targetIndex)
		// goroutine to calculate the entropy
		var wg sync.WaitGroup
		mu := &sync.Mutex{}

		for _, subset := range splits {
			wg.Add(1)
			go func(subset [][]string) {
				defer wg.Done()
				mu.Lock()
				prob := float64(len(subset)) / float64(dataRows)
				entropyAfterSplit += prob * CalculateEntropy(subset, targetName, headers)

				mu.Unlock()
			}(subset)
		}
		wg.Wait()
	} else { // calculation for the NUmericals(num,date,time)
		leftSplit, rightSplit, _ := splitByNumeric(data, targetIndex, headers)
		total := float64(len(data))

		if len(leftSplit) > 0 {
			entropyAfterSplit += (float64(len(leftSplit)) / total) * CalculateEntropy(leftSplit, targetName, headers)
		}
		if len(rightSplit) > 0 {
			entropyAfterSplit += (float64(len(rightSplit)) / total) * CalculateEntropy(rightSplit, targetName, headers)
		}

	}
	return entropyBeforeSplit - entropyAfterSplit
}
