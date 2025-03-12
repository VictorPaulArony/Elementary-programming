package utils

// function to compute the gain ratio for the C4.5 to determine the attribute values
func GainRatio(data *DataSet, targetIndex, targetColumn int) float64 {
	infoGain := CalculateInfoGain(data, targetIndex, targetColumn)
	infoSplit := SplitInformation(data.Data, targetIndex)

	if infoSplit == 0 {
		return 0
	}
	return infoGain / infoSplit
}
