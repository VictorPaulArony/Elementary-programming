package utils

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

// Nodes representing the decision tree
type Node struct {
	Feature    string           `json:"feature,omitempty"`
	Threshold  float64          `json:"threshold,omitempty"`
	Lable      string           `json:"lable,omitempty"`
	Children   map[string]*Node `json:"children,omitempty"`
	IsLeaf     bool             `json:"is_leaf"`
	SplitValue float64          `json:"SplitValue,omitempty"` // the media split value
	IsNumeric  bool             `json:"is_numeric,omitempty"`
	DataType   string           `json:"data_type,omitempty"`
}

// function to build the decision tree recursively
func DecisionTree(data [][]string, attributes []string, headers []string, targetName string) *Node {
	// targetIndex := findColumnIndex(headers, targetName)
	// case1: no data
	if len(data) == 0 {
		return &Node{IsLeaf: true, Lable: ""}
	}

	// case2: all samples are of one class/lable
	class, isPure := CheckPureClass(data, targetName, headers)
	if isPure {
		return &Node{IsLeaf: true, Lable: class}
	}

	// case3: check if attributes list is empty
	if len(attributes) == 0 {
		return &Node{IsLeaf: true, Lable: CommonClassLable(data, targetName, headers)}
	}

	// Find the best attribute to split on
	bestAttr, bestScore := BestLable(data, attributes, targetName, headers)
	// println(bestScore)
	// case4: if no good attribute found, return a leaf node
	if bestScore <= 0 {
		return &Node{IsLeaf: true, Lable: CommonClassLable(data, targetName, headers)}
	}

	bestFeature := headers[bestAttr]
	node := &Node{
		Feature:  bestFeature,
		IsLeaf:   false,
		Children: make(map[string]*Node),
		DataType: "",
	}

	dataType := DetermineDataType(data, headers)
	// Handle the split differently based on attribute type
	if dataType == "continuous" {

		// For numerical attributes, find the best threshold
		// This is a simplified version - real implementation would be more complex
		node.IsNumeric = true
		threshold, splitData := findBestThreshold(data, bestFeature, targetName, headers)
		node.Threshold = threshold

		// Create remaining attributes list (excluding the best attribute)
		remainingAttrs := filterAttributes(attributes, bestFeature)

		// Less than threshold
		node.Children["less"] = DecisionTree(splitData["less"], remainingAttrs, headers, targetName)

		// Greater than or equal to threshold
		node.Children["greater"] = DecisionTree(splitData["greater"], remainingAttrs, headers, targetName)

	} else {
		// For categorical attributes, create a child for each value
		valueGroups := make(map[string][][]string)
		attrIdx := FindColumnIndex(headers, bestFeature)

		for _, row := range data {
			attrValue := row[attrIdx]
			valueGroups[attrValue] = append(valueGroups[attrValue], row)
		}

		// Create remaining attributes list (excluding the best attribute)
		remainingAttrs := filterAttributes(attributes, bestFeature)

		// Recursively build subtrees
		for value, subset := range valueGroups {
			node.Children[value] = DecisionTree(subset, remainingAttrs, headers, targetName)
		}
	}

	return node
}

// Helper function to filter out the used attribute
func filterAttributes(attributes []string, bestFeature string) []string {
	remainingAttrs := []string{}
	for _, attr := range attributes {
		if attr != bestFeature {
			remainingAttrs = append(remainingAttrs, attr)
		}
	}
	return remainingAttrs
}

// findBestThreshold finds the optimal threshold for a numerical attribute
func findBestThreshold(data [][]string, attrName string, targetName string, headers []string) (float64, map[string][][]string) {
	attrIdx := FindColumnIndex(headers, attrName)

	// Extract and sort unique values
	values := []float64{}
	seen := make(map[float64]bool)

	for _, row := range data {
		val, err := strconv.ParseFloat(row[attrIdx], 64)
		if err == nil && !seen[val] {
			values = append(values, val)
			seen[val] = true
		}
	}

	sort.Float64s(values)

	// Try different thresholds (midpoints between consecutive values)
	bestThreshold := 0.0
	bestGainRatio := -1.0

	for i := 0; i < len(values)-1; i++ {
		threshold := (values[i] + values[i+1]) / 2

		// Split data based on threshold
		lessThan := [][]string{}
		greaterEqual := [][]string{}

		for _, row := range data {
			val, err := strconv.ParseFloat(row[attrIdx], 64)
			if err == nil {
				if val < threshold {
					lessThan = append(lessThan, row)
				} else {
					greaterEqual = append(greaterEqual, row)
				}
			}
		}

		// Calculate gain ratio for this threshold
		// This is a simplified calculation - real implementation would be more complex
		totalEntropy := CalculateEntropy(data, targetName, headers)
		lessEntropy := CalculateEntropy(lessThan, targetName, headers)
		greaterEntropy := CalculateEntropy(greaterEqual, targetName, headers)

		lessWeight := float64(len(lessThan)) / float64(len(data))
		greaterWeight := float64(len(greaterEqual)) / float64(len(data))

		infoGain := totalEntropy - (lessWeight*lessEntropy + greaterWeight*greaterEntropy)

		// Calculate split information
		splitInfo := -(lessWeight*math.Log2(lessWeight) + greaterWeight*math.Log2(greaterWeight))

		if splitInfo > 0 {
			gainRatio := infoGain / splitInfo
			if gainRatio > bestGainRatio {
				bestGainRatio = gainRatio
				bestThreshold = threshold
			}
		}
	}

	// Split data based on best threshold
	splitData := make(map[string][][]string)
	splitData["less"] = [][]string{}
	splitData["greater"] = [][]string{}

	for _, row := range data {
		val, err := strconv.ParseFloat(row[attrIdx], 64)
		if err == nil {
			if val < bestThreshold {
				splitData["less"] = append(splitData["less"], row)
			} else {
				splitData["greater"] = append(splitData["greater"], row)
			}
		}
	}

	return bestThreshold, splitData
}

// Classify a new instance using the decision tree
func Classify(node *Node, instance []string, headers []string) string {
	if node.IsLeaf {
		return node.Lable
	}

	attrIdx := FindColumnIndex(headers, node.Feature)
	attrValue := instance[attrIdx]

	if node.IsNumeric {
		// Handle numerical attributes
		val, err := strconv.ParseFloat(attrValue, 64)
		if err != nil {
			// Handle parsing error - return most common class
			return node.Lable
		}

		if val < node.Threshold {
			return Classify(node.Children["less"], instance, headers)
		} else {
			return Classify(node.Children["greater"], instance, headers)
		}
	} else {
		// Handle categorical attributes
		childNode, exists := node.Children[attrValue]
		if !exists {
			// If no matching branch, return the most common class
			return node.Lable
		}
		return Classify(childNode, instance, headers)
	}
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

	// fmt.Println("Decision tree saved to:", filePath)
}
