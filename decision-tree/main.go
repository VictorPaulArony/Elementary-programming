package main

import (
	"fmt"

	"decisision-tree/utils"
)

// function main
func main() {
	fmt.Println("This is the main function")

	dataset := utils.ReadingFIle("./training-data/data.csv", "PlayTennis")

	datatype := utils.DetermineDataType(dataset.Data, dataset.Headers)

	fmt.Println(dataset.Target)

	fmt.Println(datatype)

	tree := utils.DecisionTree(dataset.Data, dataset.Target, dataset.Headers)
	

	utils.SavePrediction("model.json", tree)
}
