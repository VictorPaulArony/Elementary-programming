package utils

// function to calculate the information gain and the gain ratio
//
//	Information Gain = Entropy(parent) - (Weighted Average) * Entropy(child)
func CalculateInfoGain(data *DataSet, targetIndex, targetColumn int) float64 {
	entropyBeforeSplit := CalcutateEntropy(data, targetIndex)
	splits := SplitData(data.Data, targetColumn)

	entropyAfterSplit := 0.0
	dataRows := len(data.Data)

	for _, subset := range splits {
		subsetDataSet := &DataSet{
			Headers: data.Headers,
			Data:    subset,
		}
		prob := float64(len(subset)) / float64(dataRows)
		entropyAfterSplit += prob * CalcutateEntropy(subsetDataSet, targetIndex)

	}

	return entropyBeforeSplit - entropyAfterSplit
}

// function to split dataset int subsets based on an column/attributes
func SplitData(data [][]string, columnIndex int) map[string][][]string {
	// a map to hold the splits, where the key is the column value,
	// and the value is a slice of rows that have that column value.
	splits := make(map[string][][]string)

	for _, row := range data {
		val := row[columnIndex]
		splits[val] = append(splits[val], row)
	}
	return splits
}
