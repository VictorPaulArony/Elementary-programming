package utils

// function to compute the gain ratio for the C4.5 to determine the attribute values
func GainRatio(data [][]string, attrName string, targetName string, headers []string) float64 {
	infoGain := CalculateInfoGain(data, attrName, targetName, headers)
	infoSplit := SplitInformation(data, attrIndex)
	if infoSplit == 0 {
		return 0.0
	}
	// println("hello",infoGain / infoSplit)
	return infoGain / infoSplit
}

// to find the best lable/ class to split from
func BestLable(data [][]string, attributes []string, targetName string, headers []string) (int, float64) {
	targetIndex := FindColumnIndex(headers, targetName)
	bestLable := -1
	bestScore := 0.0

	for lableINdex := range headers {
		if lableINdex == targetIndex {
			continue
		}
		score := GainRatio(data, targetName, targetName, headers)
		if score > bestScore {
			bestScore = score
			bestLable = lableINdex
		}
	}
	
	return bestLable, bestScore
}

// function to determine the most common class lable
func CommonClassLable(data [][]string, targetName string, headers []string) string {
	targetIndex := FindColumnIndex(headers, targetName)
	classCount := make(map[string]int)
	mostClass := ""
	maxCount := 0

	for _, row := range data {
		classCount[row[targetIndex]]++
	}

	for class, count := range classCount {
		if count > maxCount {
			maxCount = count
			mostClass = class
		}
	}
	return mostClass
}

// function to check if all the samples belong to the same calss
func CheckPureClass(data [][]string, targetName string, headers []string) (string, bool) {
	if len(data) == 0 {
		return "", true
	}

	targetIndex := FindColumnIndex(headers, targetName)
	firstClass := data[0][targetIndex]

	for _, row := range data {
		if row[targetIndex] != firstClass {
			return "", false
		}
	}
	return firstClass, true
}
