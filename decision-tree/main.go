package main

import (
	"fmt"

	"decisision-tree/utils"
)

// function main
func main() {
    fmt.Println("This is the main function")
    dataset := utils.ReadingFIle("./training-data/heart_2022_with_nans.csv", "CovidPos")
    // dataset := utils.ReadingFIle("./training-data/data.csv", "PlayTennis")
    targetIndex := utils.FindColumnIndex(dataset.Headers, dataset.Target)
    fmt.Println(dataset.Headers)
    datatype := utils.DetermineDataType(dataset.Data, targetIndex)
    // fmt.Println(dataset.Data)
    fmt.Println(datatype)
    // println(targetIndex, dataset.Target)
    entropy := utils.CalculateEntropy(dataset.Data, targetIndex)
    fmt.Println(entropy)
    
    // Fix the attributes list creation - you're currently appending the entire headers slice multiple times
    attributes := []string{}
    for i, header := range dataset.Headers {
        if i == targetIndex {
            continue
        }
        attributes = append(attributes, header)
    }
    
    // Call DecisionTree with the correct attributes
    tree := utils.ParallelDecisionTree(dataset.Data, attributes, dataset.Headers, dataset.Target)
    utils.SavePrediction("model.json", tree)
    
    // Update GainRatio call to include both the column index and target index
    // Choose a specific column for analysis, or loop through all columns
    // For example, for the first attribute:
    if len(attributes) > 0 {
        columnIndex := utils.FindColumnIndex(dataset.Headers, attributes[0])
        gain := utils.GainRatio(dataset.Data, columnIndex, targetIndex)
        fmt.Printf("Gain ratio for %s: %f\n", attributes[0], gain)
    }
    
    // Or calculate gain ratio for all attributes:
    // fmt.Println("Gain ratios for all attributes:")
    // for _, attr := range attributes {
    //     columnIndex := utils.FindColumnIndex(dataset.Headers, attr)
    //     gain := utils.GainRatio(dataset.Data, columnIndex, targetIndex)
    //     fmt.Printf("%s: %f\n", attr, gain)
    // }
}
