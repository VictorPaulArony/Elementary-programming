package utils

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// function to split dataset int subsets based on an column/attributes

func splitDataCategoricalSequential(data [][]string, columnIndex int) map[string][][]string {
	// a map to hold the splits, where the key is the column value,
	// and the value is a slice of rows that have that column value.
	splits := make(map[string][][]string)

	for _, row := range data {
		if len(row) > columnIndex { // Ensure index is within bounds
			val := row[columnIndex]
			splits[val] = append(splits[val], row)
		}
	}
	return splits
}

// function to split data by numeri, date and time using median
func SequentialSplitByNumeric(data [][]string, columnIndex int) ([][]string, [][]string, float64) {
	values := []float64{}
	leftSplit, rightSplit := [][]string{}, [][]string{}

	// Extract numeric values
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

	// Split data based on median
	for _, row := range data {
		if len(row) <= columnIndex {
			continue
		}
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

	fmt.Println(leftSplit, " , ", rightSplit)
	return leftSplit, rightSplit, median
}

// Parallel function to split data by numeric, date, and time using median
func SplitByNumericParallel(data [][]string, columnIndex int) ([][]string, [][]string, float64) {
	values := []float64{}
	leftSplit, rightSplit := [][]string{}, [][]string{}
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect shared access

	// Extract numeric values in parallel
	for _, row := range data {
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
			if len(row) <= columnIndex {
				return
			}

			num, err := strconv.ParseFloat(row[columnIndex], 64)
			if err != nil {
				date, err := time.Parse("2006-01-02", row[columnIndex])
				if err != nil {
					mu.Lock()
					leftSplit = append(leftSplit, row) // Append to leftSplit on error
					mu.Unlock()
					return
				}
				num = float64(date.Unix())
			}

			mu.Lock()
			values = append(values, num) // Append the parsed numeric value
			mu.Unlock()
		}(row)
	}

	wg.Wait() // Wait for all goroutines to finish

	if len(values) == 0 {
		return data, [][]string{}, 0.0 // Return original data if no valid values
	}

	sort.Float64s(values) // For median computation
	lenValues := len(values)
	median := 0.0

	if lenValues%2 == 0 { // Even data: take the average of the two middle values
		median = (values[lenValues/2-1] + values[lenValues/2]) / 2.0
	} else { // Odd length data: take the middle value
		median = values[lenValues/2]
	}

	// Split data based on median
	for _, row := range data {
		if len(row) <= columnIndex {
			continue
		}
		num, err := strconv.ParseFloat(row[columnIndex], 64)
		if err != nil {
			date, err := time.Parse("2006-01-02", row[columnIndex])
			if err != nil {
				leftSplit = append(leftSplit, row) // Default to left if parsing fails
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

// Switch function to choose between the two implementations
func SplitByNumeric(data [][]string, columnIndex int) ([][]string, [][]string, float64) {
	if len(data) > 10000 {
		return SplitByNumericParallel(data, columnIndex)
	}
	return SequentialSplitByNumeric(data, columnIndex)
}
