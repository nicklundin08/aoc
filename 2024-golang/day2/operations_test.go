package main

import (
	"fmt"
	"testing"
)

func TestRecursiveCall(t *testing.T) {
	testCases := []struct {
		values      []int
		expectation bool
	}{
		{
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 1, 1},
			false,
		},
		{
			[]int{51, 54, 57, 60, 61, 64, 67, 64},
			false,
		},
	}
	for _, tc := range testCases {
		result := rowPasses(tc.values, false)
		if result != tc.expectation {
			t.Errorf("Expected %t but found %t with values %v\n", tc.expectation, result, tc.values)
		}
	}
}

func TestFileWithNoDampener(t *testing.T) {
	fmt.Println("----------------------- Hello world")
	path := "./input.txt"
	result, err := readListsFromFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Total number of lines in file %d.\n", len(result))
	passCount := 0
	for _, row := range result {
		if rowPasses(row, false) {
			passCount++
		}
	}
	fmt.Printf("pass count: %d.\n", passCount)
	fmt.Println("----------------------- Goodbye world")
	if passCount != 224 {
		t.Errorf("Found %d pass count instead of 224", passCount)
	}
}

func TestFileWithDampener(t *testing.T) {
	fmt.Println("----------------------- Hello world")
	path := "./input.txt"
	result, err := readListsFromFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Total number of lines in file %d.\n", len(result))
	passCount := 0
	for _, row := range result {
		if rowPasses(row, true) {
			passCount++
		}
	}
	fmt.Printf("pass count: %d.\n", passCount)
	fmt.Println("----------------------- Goodbye world")
	if passCount != 224 {
		t.Errorf("Found %d pass count instead of 224", passCount)
	}
}
