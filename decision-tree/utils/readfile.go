package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type DataSet struct {
	Headers  []string
	Data     [][]string
	DataType map[string]string // continuous, categorical, data, timestamp
	Target   string
}

// function to read the CSV file
func ReadingFIle(fileName, targetIndex string) DataSet {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error: Missing input file ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Eroor while reading the file data: ", err)
	}

	headings := rows[0] // the first row is for the headings
	data := rows[1:]    // the raw data of the file

	return DataSet{
		Headers:  headings,
		Data:     data,
		Target:   targetIndex,
		DataType: make(map[string]string),
	}
}

// function to determine if data type is continuous, categorical, or date
func DetermineDataType(data [][]string, headers []string) string {
	res := ""

	// goroutine for big data sets >10K
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	targetIdx := -1
	for i := range headers {
		if i == targetIdx {
			continue
		}

		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			isNumerical := true // numerical until proven to get category data
			isDate := false     // date
			isTime := false     // time

			for _, row := range data {
				// Check for missing values
				if row[i] == "" {
					isNumerical = false // If there's a missing value, can't be fully numerical
					continue
				}

				// user parse as float
				if _, err := strconv.ParseFloat(row[i], 64); err == nil {
					isNumerical = true
				} else {
					isNumerical = false
				}

				// dtermine date
				if _, err := time.Parse("2006-01-02", row[i]); err == nil {
					isDate = true
				}

				// determine time
				if _, err := time.Parse(time.RFC3339, row[i]); err == nil {
					isTime = true
				}

			}

			// Assign the correct data type
			mu.Lock()
			if isNumerical {
				res = "continuous"
			} else if isDate {
				res = "date"
			} else if isTime {
				res = "time"
			} else {
				res = "categorical"
			}
			mu.Unlock()
		}(i)

	}
	wg.Wait() // wait for all goroutin to finish
	return res
}
