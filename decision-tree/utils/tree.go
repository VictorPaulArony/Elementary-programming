package utils

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
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
func ParallelDecisionTree(data [][]string, attributes []string, headers []string, targetName string) *Node {
	targetIndex := FindColumnIndex(headers, targetName)
	if targetIndex == -1 {
		return &Node{IsLeaf: true, Lable: ""}
	}

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
	bestAttr, bestScore := ParallelBestLable(data, attributes, targetName, headers)

	// case4: if no good attribute found, return a leaf node
	if bestScore <= 0 || bestAttr == -1 {
		return &Node{IsLeaf: true, Lable: CommonClassLable(data, targetName, headers)}
	}

	bestFeature := headers[bestAttr]
	dataType := DetermineDataType(data, bestAttr)

	node := &Node{
		Feature:  bestFeature,
		IsLeaf:   false,
		Children: make(map[string]*Node),
		DataType: dataType,
	}

	// Create remaining attributes list (excluding the best attribute)
	remainingAttrs := filterAttributes(attributes, bestFeature)

	// Handle the split differently based on attribute type
	if dataType == "continuous" {
		// For numerical attributes, find the best threshold
		node.IsNumeric = true
		threshold, splitData := findBestThreshold(data, bestFeature, headers)
		node.Threshold = threshold

		// Check if split was successful
		if len(splitData["less"]) == 0 || len(splitData["greater"]) == 0 {
			return &Node{IsLeaf: true, Lable: CommonClassLable(data, targetName, headers)}
		}

		// Process left and right branches in parallel
		var wg sync.WaitGroup
		wg.Add(2)

		var leftNode, rightNode *Node

		// Process left subtree in a goroutine
		go func() {
			defer wg.Done()
			leftNode = ParallelDecisionTree(splitData["less"], remainingAttrs, headers, targetName)
		}()

		// Process right subtree in a goroutine
		go func() {
			defer wg.Done()
			rightNode = ParallelDecisionTree(splitData["greater"], remainingAttrs, headers, targetName)
		}()

		// Wait for both goroutines to complete
		wg.Wait()

		// Assign the results
		node.Children["less"] = leftNode
		node.Children["greater"] = rightNode
	} else {
		// For categorical attributes, create a child for each value
		valueGroups := make(map[string][][]string)
		attrIdx := FindColumnIndex(headers, bestFeature)

		// Use a mutex to protect concurrent access to valueGroups
		var mu sync.Mutex
		var wg sync.WaitGroup

		// Process rows in parallel for large datasets
		batchSize := 1000
		if len(data) > batchSize {
			for i := 0; i < len(data); i += batchSize {
				wg.Add(1)
				end := i + batchSize
				if end > len(data) {
					end = len(data)
				}

				go func(start, end int) {
					defer wg.Done()
					localGroups := make(map[string][][]string)

					for j := start; j < end; j++ {
						row := data[j]
						if len(row) <= attrIdx {
							continue
						}
						attrValue := row[attrIdx]
						localGroups[attrValue] = append(localGroups[attrValue], row)
					}

					// Merge local results into global map
					mu.Lock()
					for value, rows := range localGroups {
						valueGroups[value] = append(valueGroups[value], rows...)
					}
					mu.Unlock()
				}(i, end)
			}
			wg.Wait()
		} else {
			// Sequential processing for small datasets
			for _, row := range data {
				if len(row) <= attrIdx {
					continue
				}
				attrValue := row[attrIdx]
				valueGroups[attrValue] = append(valueGroups[attrValue], row)
			}
		}

		// Check if split was successful
		if len(valueGroups) <= 1 {
			return &Node{IsLeaf: true, Lable: CommonClassLable(data, targetName, headers)}
		}

		// Create a channel to receive results
		type childResult struct {
			value string
			node  *Node
		}
		resultChan := make(chan childResult, len(valueGroups))

		// Launch a goroutine for each value group to build subtrees in parallel
		for value, subset := range valueGroups {
			wg.Add(1)
			go func(val string, data [][]string) {
				defer wg.Done()
				childNode := ParallelDecisionTree(data, remainingAttrs, headers, targetName)
				resultChan <- childResult{val, childNode}
			}(value, subset)
		}

		// Use a separate goroutine to collect results
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		// Collect all results from the channel
		for result := range resultChan {
			node.Children[result.value] = result.node
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
func findBestThreshold(data [][]string, featureName string, headers []string) (float64, map[string][][]string) {
	featureIndex := FindColumnIndex(headers, featureName)
	if featureIndex == -1 {
		return 0, nil
	}

	// Get all unique values for the feature
	values := make([]float64, 0)
	for _, row := range data {
		if len(row) <= featureIndex {
			continue
		}

		val, err := strconv.ParseFloat(row[featureIndex], 64)
		if err != nil {
			// Try to parse as date
			date, err := time.Parse("2006-01-02", row[featureIndex])
			if err != nil {
				continue
			}
			val = float64(date.Unix())
		}

		values = append(values, val)
	}

	if len(values) == 0 {
		return 0, map[string][][]string{"less": data, "greater": {}}
	}

	// Sort values
	sort.Float64s(values)

	// Find median
	var threshold float64
	if len(values)%2 == 0 {
		threshold = (values[len(values)/2-1] + values[len(values)/2]) / 2
	} else {
		threshold = values[len(values)/2]
	}

	// Split data based on threshold
	splitData := map[string][][]string{
		"less":    {},
		"greater": {},
	}

	for _, row := range data {
		if len(row) <= featureIndex {
			continue
		}

		val, err := strconv.ParseFloat(row[featureIndex], 64)
		if err != nil {
			// Try to parse as date
			date, err := time.Parse("2006-01-02", row[featureIndex])
			if err != nil {
				splitData["less"] = append(splitData["less"], row)
				continue
			}
			val = float64(date.Unix())
		}

		if val < threshold {
			splitData["less"] = append(splitData["less"], row)
		} else {
			splitData["greater"] = append(splitData["greater"], row)
		}
	}

	return threshold, splitData
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
