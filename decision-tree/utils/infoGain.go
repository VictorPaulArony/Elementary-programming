package utils

import (
	"sync"
)

// function to calculate the information gain
// information Gain = Entropy(parent) - (Weighted Average) * Entropy(child)
func CalculateInfoGain(data [][]string, targetIndex, targetColumn int) float64 {
	entropyBeforeSplit := CalculateEntropy(data, targetIndex)


	colType := DetermineDataType(data, targetColumn)

	entropyAfterSplit := 0.0
	dataRows := len(data)

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
				entropyAfterSplit += prob * CalculateEntropy(subset, targetIndex)

				mu.Unlock()
			}(subset)
		}
		wg.Wait()
	} else { // calculation for the NUmericals(num,date,time)
		leftSplit, rightSplit, _ := splitByNumeric(data, targetColumn)
		total := float64(len(data))

		entropyAfterSplit = (float64(len(leftSplit))/total)*CalculateEntropy(leftSplit, targetIndex) +
			(float64(len(rightSplit))/total)*CalculateEntropy(rightSplit, targetIndex)
	}

	return entropyBeforeSplit - entropyAfterSplit
}
