package main

import (
	"errors"
	"fmt"
	"sort"
)

func abs(x int) int {
	if x > 0 {
		return x
	}

	return 0 - x
}

func sum(list []int) int{
  result := 0
	for i := range list {
    result += list[i]
  }
  return result

}

func printLists(l1 []int, l2 []int) {
	fmt.Println("list1")
	fmt.Println(l2)
	fmt.Println("list2")
	fmt.Println(l1)
}

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

func main() {
	fmt.Println("hello world")
	l1 := []int{3, 4, 2, 1, 3, 3}
	l2 := []int{4, 3, 5, 3, 9, 3}
	fmt.Println("------------------- before sorting")
	printLists(l1, l2)

	fmt.Println("------------------- after sorting")
	sort.Ints(l1)
	sort.Ints(l2)
	printLists(l1, l2)

	result, err := diff(l1, l2)
	if err != nil {
		fmt.Println("an error occured!")
		return
	}
	fmt.Println("------------------- difference list")
	fmt.Println(result)

	fmt.Println("------------------- total distance")
  sum := sum(result)
	fmt.Println(sum)
}
