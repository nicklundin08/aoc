package main

import (
	"errors"
)

func abs(x int) int {
	if x > 0 {
		return x
	}

	return 0 - x
}

func sum(list []int) int {
	result := 0
	for i := range list {
		result += list[i]
	}
	return result
}

// Given two slices of equal size, returns a new slice of the difference between the two slices values
// e.g. diff({100,200,400}, {100, 100, 100}) -> {0, 100, 300}, nil
func diff(l1 []int, l2 []int) ([]int, error) {
	if len(l1) != len(l2) {
		return nil, errors.New("Slices must be same len!")
	}

	result := make([]int, len(l1))
	for i := range l1 {
		x1 := l1[i]
		x2 := l2[i]
		x3 := x1 - x2
		result[i] = abs(x3)
	}
	return result, nil
}

// Computes a similarity score for two lists
func similarityScore(l1 []int, l2 []int) (int, error) {
	//This is the multiplier value for the right hand list.
	//e.g an entry {5: 2}, would signal that the value "5" occurs "2" times in the right hand list
  rightHandSideMultiplier := make(map[int]int)
  leftHandSideMultiplier := make(map[int]int)

	for _, value := range l1 {
		currentMultiplier := leftHandSideMultiplier[value]
		currentMultiplier += 1
		leftHandSideMultiplier[value] = currentMultiplier
	}

	for _, value := range l2 {
		currentMultiplier := rightHandSideMultiplier[value]
		currentMultiplier += 1
		rightHandSideMultiplier[value] = currentMultiplier
	}

	similarityScore := 0

	for key, lValue := range leftHandSideMultiplier {
    rValue := rightHandSideMultiplier[key]
    result := key * lValue * rValue
		similarityScore += result
	}
	return similarityScore, nil
}
