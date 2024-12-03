package main

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		reasoning   string
		values      []int
		expectation bool
	}{
		{
			values:      []int{1, 2, 3},
			expectation: true,
			reasoning:   "Should pass without removing any level",
		},
		{
			values:      []int{7, 1, 2, 3, 4},
			expectation: true,
			reasoning:   "Should pass if you remove the first level",
		},
		{
			values:      []int{0, 1, 2, 3, 10000},
			expectation: true,
			reasoning:   "Should pass if you remove the last level",
		},
		{
			values:      []int{0, 5, 1, 300, 10000},
			expectation: false,
			reasoning:   "Should fail.",
		},
		{
			values:      []int{7, 6, 4, 2, 1},
			expectation: true,
			reasoning:   "Safe without removing any values",
		},
		{
			values:      []int{1, 2, 7, 8, 9},
			expectation: false,
			reasoning:   "Unsafe regardless",
		},
		{
			values:      []int{9, 7, 6, 2, 1},
			expectation: false,
			reasoning:   "Unsafe regardless",
		},
		{
			values:      []int{1, 3, 2, 4, 5},
			expectation: true,
			reasoning:   "Safe by removing the 2nd value",
		},
		{
			values:      []int{8, 6, 4, 4, 1},
			expectation: true,
			reasoning:   "Save by removing the 3rd value",
		},
		{
			values:      []int{1, 3, 6, 7, 9},
			expectation: true,
			reasoning:   "Safe",
		},
	}
	// testCases = testCases[5:7]
	for _, tc := range testCases {
		fmt.Println("-------------")
		result := check2WithSkips(tc.values)
		if result != tc.expectation {
			t.Errorf("Expected %t but found %t with values %v, reasoning %s.\n", tc.expectation, result, tc.values, tc.reasoning)
		}
	}
}

func TestFileWithNoDampener(t *testing.T) {
	// fmt.Println("----------------------- Hello world")
	path := "./input.txt"
	result, err := readListsFromFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	// fmt.Printf("Total number of lines in file %d.\n", len(result))
	passCount := 0
	for _, row := range result {
		if check2WithSkips(row) {
			passCount++
		}
	}
	// fmt.Printf("pass count: %d.\n", passCount)
	// fmt.Println("----------------------- Goodbye world")
	if passCount != 224 {
		t.Errorf("Found %d pass count instead of 224", passCount)
	}
}

func TestFileWithDampener(t *testing.T) {
	// fmt.Println("----------------------- Hello world")
	path := "./input.txt"
	result, err := readListsFromFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	// fmt.Printf("Total number of lines in file %d.\n", len(result))
	passCount := 0
	for _, row := range result {
		if check2WithSkips(row) {
			passCount++
		}
	}
	// fmt.Printf("pass count: %d.\n", passCount)
	// fmt.Println("----------------------- Goodbye world")
	if passCount <= 224 {
		t.Errorf("Should be greater than 224. %d", passCount)
	}
	if passCount > 350 {
		t.Errorf("Should be less than 350. %d", passCount)
	}
}
