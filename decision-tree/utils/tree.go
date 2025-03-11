package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Nodes representing the decision tree
type Node struct {
	Feature   string           `json:"feature,omitempty"`
	Threshold float64          `json:"threshold,omitempty"`
	Lable     string           `json:"lable,omitempty"`
	Children  map[string]*Node `json:"children,omitempty"`
	IsLeaf    bool             `json:"is_leaf"`
}

// function to save the predictions to the JSON file
func SavePrediction(filePath string, tree *Node) {
	data, err := json.MarshalIndent(tree, "", " ")
	if err != nil {
		log.Println("ERROR: Unable to Marshal data of that formart: ", err)
		return
	}

	if err = os.WriteFile(filePath, data, 0o644); err != nil {
		log.Println("Error: Output path not	specified: ", err)
		return
	}

	fmt.Println("Decision tree saved to:", filePath)
}
