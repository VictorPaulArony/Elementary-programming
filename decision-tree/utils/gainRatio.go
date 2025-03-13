package utils

// function to compute the gain ratio for the C4.5 to determine the attribute values
func GainRatio(data [][]string,attrName string, targetName string, headers []string) float64 {
	infoGain := CalculateInfoGain(data, attrName, targetName, headers)
	infoSplit := SplitInformation(data, attrName, targetName, headers)

	if infoSplit == 0 {
		return 0.0
	}
	return infoGain / infoSplit
}

// to find the best lable/ class to split from
func BestLable(data [][]string, attributes []string, targetName string, headers []string) (string, float64) {
	bestLable := ""
	bestScore := -1.0

	for _, lablecolumn := range headers {
		score := GainRatio(data, lablecolumn, targetName, headers)
		if score > bestScore {
			bestScore = score
			bestLable = lablecolumn
		}
	}
	return bestLable, bestScore
}

// function to determine the most common class lable
func CommonClassLable(data [][]string, targetIndex int) string {
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
func CheckPureClass(data [][]string, targetIndex int) (string, bool) {
	if len(data) == 0 {
		return "", true
	}

	firstClass := data[0][targetIndex]
	for _, row := range data {
		if row[targetIndex] != firstClass {
			return "", false
		}
	}
	return firstClass, true
}
