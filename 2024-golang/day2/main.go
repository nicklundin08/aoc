package main

import (
	"fmt"
)

func main() {
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
}
