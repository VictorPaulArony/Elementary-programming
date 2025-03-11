package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type DataSet struct {
	Headers  []string
	Data     [][]string
	DataType map[string]string // continuous or categorical
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

// function to determine if data type is continous or categorical
func DetermineDataType(data *DataSet) {
	for i, header := range data.Headers {
		if header == data.Target {
			continue
		}

		isContinuous := true // numerical
		for _, row := range data.Data {
			// user parse as float
			_, err := strconv.ParseFloat(row[i], 64)
			if err != nil {
				isContinuous = false
				break
			}
		}

		if isContinuous {
			data.DataType[header] = "continuous"
		} else {
			data.DataType[header] = "categorical"
		}

	}
}
