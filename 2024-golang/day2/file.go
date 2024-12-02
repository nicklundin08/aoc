package main

import (
	"os"
	"strconv"
	"strings"
)

// Reads a file into an [][]int. Assumes the file has a specific format.
// 1000 6 10 302
// 109 945 6 10 302

// Would result in:
// [][]int{{1000, 6, 10, 302},{109, 945, 6, 10, 32}}
func readListsFromFile(path string) ([][]int, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fileStr := string(fileBytes)
	lines := strings.Split(fileStr, "\n")

	result := make([][]int, len(lines))

	for _, line := range lines {
		values := strings.Fields(line)
		currentIntArray := make([]int, len(values))
		for index, value := range values {
			intVal, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			currentIntArray[index] = intVal
		}
		result = append(result, currentIntArray)
	}
	return result, nil
}
