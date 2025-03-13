package utils

import (
	"encoding/json"
	"log"
	"os"
)

// Nodes representing the decision tree
type Node struct {
	Index      int              `json:"index,omitempty"`
	Feature    string           `json:"feature,omitempty"`
	Threshold  float64          `json:"threshold,omitempty"`
	Lable      string           `json:"lable,omitempty"`
	Children   map[string]*Node `json:"children,omitempty"`
	IsLeaf     bool             `json:"is_leaf"`
	SplitValue float64          `json:"SplitValue,omitempty"` // the media split value
	IsNumeric  bool             `json:"is_numeric,omitempty"`
}

// function to build the decision tree recursively
// func DecisionTree(data *DataSet, attributes []string ) *Node {
// 	// case1: no data
// 	if len(data.Data) == 0 {
// 		return &Node{IsLeaf: true, Lable: ""}
// 	}

// 	// case2: all samples are of one class/lable
// 	class, isPure := CheckPureClass(data, targetIndex, headers)
// 	if isPure {
// 		return &Node{IsLeaf: true, Lable: class}
// 	}

// 	// Find the best attribute to split on
// 	bestAttr, bestScore := BestLable(data, targetIndex, headers)

// 	// // If no good attribute found, return a leaf node
// 	if bestAttr == "-1" || bestScore <= 0 {
// 		return &Node{IsLeaf: true, Lable: CommonClassLable(data, targetIndex, headers)}
// 	}

// 	// Get the index of the best attribute
// 	bestAttrIndex := -1
// 	for i, attr := range headers {
// 		if attr == bestAttr {
// 			bestAttrIndex = i
// 			break
// 		}
// 	}

// 	// Check if attribute is numeric
// 	isNumeric := false
// 	if _, err := strconv.ParseFloat(data[0][bestAttrIndex], 64); err == nil {
// 		isNumeric = true
// 	}
// 	node := &Node{
// 		Index:     bestAttrIndex,
// 		IsLeaf:    false,
// 		IsNumeric: isNumeric,
// 		Children:  make(map[string]*Node),
// 	}

// 	// Create remaining attributes list (remove the best attribute)
// 	remainingAttrs := make([]string, 0)
// 	for _, attr := range headers {
// 		if attr != headers[bestAttrIndex] {
// 			remainingAttrs = append(remainingAttrs, attr)
// 		}
// 	}

// 	if isNumeric {
// 		// Handle numeric attribute split
// 		leftSplit, rightSplit, median := splitByNumeric(data, bestAttr, headers)
// 		node.SplitValue = median

// 		// Build left subtree
// 		node.Children["<= "+fmt.Sprintf("%.2f", median)] = DecisionTree(leftSplit, targetIndex, remainingAttrs)

// 		// Build right subtree
// 		node.Children["> "+fmt.Sprintf("%.2f", median)] = DecisionTree(rightSplit, targetIndex, remainingAttrs)
// 	} else {
// 		// Handle categorical attribute split
// 		splits := SplitDataCategorical(data, bestAttr)

// 		// For each value of the attribute, create a new branch
// 		for value, subset := range splits {
// 			node.Children[value] = DecisionTree(subset, targetIndex, remainingAttrs)
// 		}
// 	}

// 	return node
// }

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

	// fmt.Println("Decision tree saved to:", filePath)
}
