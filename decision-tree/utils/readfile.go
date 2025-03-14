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
func DetermineDataType(data [][]string, columnIndex int) string {
	if len(data) == 0 {
		return "categorical" // Default to categorical for empty data
	}

	// Check if the column contains numerical values
	for _, row := range data {
		if len(row) <= columnIndex {
			continue
		}

		// Try to parse as float
		_, err := strconv.ParseFloat(row[columnIndex], 64)
		if err == nil {
			return "continuous"
		}

		// Try to parse as date
		_, err = time.Parse("2006-01-02", row[columnIndex])
		if err == nil {
			return "continuous"
		}
	}

	return "categorical"
}
