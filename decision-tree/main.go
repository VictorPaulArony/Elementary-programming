package main

import (
	"fmt"

	"decisision-tree/utils"
)

// function main
func main() {
	fmt.Println("This is the main function")

	dataset := utils.ReadingFIle("./training-data/data3.csv", "")

	datatype := utils.DetermineDataType(dataset.Data, 0)

	fmt.Println(datatype)

	fmt.Println(dataset.Headers)
}
