package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

//Reads a file into 2 lists. Assumes the file has a specific format.
//1000     309
//2028     1191
//5440     2

// Would result in:
// int []{1000, 2028, 5440}
// int []{309, 1191, 2}
// nil
func readListsFromFile(path string) ([]int, []int, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	fileStr := string(fileBytes)
	values := strings.Fields(fileStr)
	if len(values) % 2 != 0 {
		return nil, nil, errors.New("Expected an array with an even length!!")
	}

  lengthToAllocate := len(values) / 2
	l1 := make([]int, lengthToAllocate)
	l2 := make([]int, lengthToAllocate)
	for index, value := range values {
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return nil, nil, err
		}
		if index%2 == 0 {
			l1 = append(l1, intVal)
		} else {
			l2 = append(l2, intVal)
		}
	}
	return l1, l2, nil
}
