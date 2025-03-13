package utils

import (
	"sort"
	"strconv"
	"time"
)

// function to split dataset int subsets based on an column/attributes
func SplitDataCategorical(data [][]string, columnIndex int) map[string][][]string {
	// a map to hold the splits, where the key is the column value,
	// and the value is a slice of rows that have that column value.
	splits := make(map[string][][]string)

	for _, row := range data {
		val := row[columnIndex]
		splits[val] = append(splits[val], row)
	}
	return splits
}

// function to split data by numeri, date and time using median
func splitByNumeric(data [][]string, columnIndex int) ([][]string, [][]string, float64) {
	values := []float64{}
	// valueToRow := make(map[float64][]string)
	leftSplit, rightSplit := [][]string{}, [][]string{}

	for _, row := range data {
		num, err := strconv.ParseFloat(row[columnIndex], 64)
		if err != nil {
			date, err := time.Parse("2006-01-02", row[columnIndex])
			if err != nil {
				continue // skip if no number or date
			}
			// convert date/time to numerical val
			num = float64(date.Unix())
		}
		values = append(values, num)
		// valueToRow[num] = row
	}

	if len(values) == 0 {
		return data, [][]string{}, 0.0 // return original data if no valid values
	}

	sort.Float64s(values) // for median computation

	lenValues := len(values)
	median := 0.0

	if lenValues%2 == 0 { // even data take the average of the two middle values
		median = (values[lenValues/2-1] + values[lenValues/2]) / 2.0
	} else { // odd length data take the middle value
		median = values[lenValues/2]
	}

	for _, row := range data {
		num, err := strconv.ParseFloat(row[columnIndex], 64)
		if err != nil {
			date, err := time.Parse("2006-01-02", row[columnIndex])
			if err != nil {
				leftSplit = append(leftSplit, row) // default to left if parsing fails
				continue
			}
			num = float64(date.Unix())
		}

		if num <= median {
			leftSplit = append(leftSplit, row)
		} else {
			rightSplit = append(rightSplit, row)
		}
	}
	return leftSplit, rightSplit, median
}
