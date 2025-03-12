package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type DataSet struct {
	Headers  []string
	Data     [][]string
	DataType map[string]string // continuous, categorical, data, timestamp
	Target   string
}

// function to read the CSV file
func ReadingFIle(fileName, targetColumn string) DataSet {
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
		Target:   targetColumn,
		DataType: make(map[string]string),
	}
}

// function to determine if data type is continuous, categorical, or date
func DetermineDataType(data *DataSet) {
	for i, header := range data.Headers {
		if header == data.Target {
			continue
		}

		isNumerical := true // numerical until proven to get category data
		isDate := false     // date
		isTime := false     // time
		for _, row := range data.Data {

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

		if isNumerical {
			data.DataType[header] = "continuous"
		} else if isDate {
			data.DataType[header] = "date"
		} else if isTime {
			data.DataType[header] = "time"
		} else {
			data.DataType[header] = "categorical"
		}

	}
}
