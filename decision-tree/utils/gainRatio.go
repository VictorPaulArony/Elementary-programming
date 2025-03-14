package utils

import "sync"

// function to compute the gain ratio for the C4.5 to determine the attribute values
func GainRatio(data [][]string, columnIndex int, targetIndex int) float64 {
	infoGain := CalculateInfoGain(data, columnIndex, targetIndex)
	infoSplit := SplitInformation(data, columnIndex)
	if infoSplit == 0 {
		return 0.0
	}
	// println("hello",infoGain / infoSplit)
	return infoGain / infoSplit
}

// to find the best lable/ class to split from
func ParallelBestLable(data [][]string, attributes []string, targetName string, headers []string) (int, float64) {
	targetIndex := FindColumnIndex(headers, targetName)
	if targetIndex == -1 {
		return -1, 0.0
	}

	// Use a mutex to protect access to shared variables
	var mu sync.Mutex
	var wg sync.WaitGroup

	bestScore := 0.0
	bestAttrIndex := -1

	// Process attributes in parallel
	for _, attr := range attributes {
		wg.Add(1)
		go func(attribute string) {
			defer wg.Done()

			attrIndex := FindColumnIndex(headers, attribute)
			if attrIndex == -1 {
				return
			}

			// Calculate gain ratio for this attribute
			score := GainRatio(data, attrIndex, targetIndex)

			// Update best score if this is better
			mu.Lock()
			if score > bestScore {
				bestScore = score
				bestAttrIndex = attrIndex
			}
			mu.Unlock()
		}(attr)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return bestAttrIndex, bestScore
}

// function to determine the most common class lable
func CommonClassLable(data [][]string, targetName string, headers []string) string {
	targetIndex := FindColumnIndex(headers, targetName)
	if targetIndex == -1 {
		return ""
	}

	classCounts := make(map[string]int)
	for _, row := range data {
		if len(row) > targetIndex {
			classCounts[row[targetIndex]]++
		}
	}

	maxCount := 0
	maxClass := ""
	for class, count := range classCounts {
		if count > maxCount {
			maxCount = count
			maxClass = class
		}
	}

	return maxClass
}

// function to check if all the samples belong to the same calss
func CheckPureClass(data [][]string, targetName string, headers []string) (string, bool) {
	targetIndex := FindColumnIndex(headers, targetName)
	if len(data) == 0 || targetIndex == -1 {
		return "", false
	}

	firstClass := data[0][targetIndex]
	for _, row := range data {
		if row[targetIndex] != firstClass {
			return "", false
		}
	}
	return firstClass, true
}
