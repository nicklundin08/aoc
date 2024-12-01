package main

import (
	"fmt"
	"sort"
)

func main() {
	//pt1Main()
	pt2Main()
}

func printLists(l1 []int, l2 []int) {
	fmt.Println("list1")
	fmt.Println(l2)
	fmt.Println("list2")
	fmt.Println(l1)
}

func pt1Main() {
	fmt.Println("hello world")
	path := "./input.txt"
	l1, l2, err := readListsFromFile(path)
	if err != nil {
		fmt.Println("an error occured!")
		return
	}

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

func pt2Main() {
	fmt.Println("hello world")
	path := "./input.txt"
	l1, l2, err := readListsFromFile(path)
	if err != nil {
		fmt.Println("an error occured!")
		return
	}

	fmt.Println("------------------- before sorting")
	printLists(l1, l2)

	fmt.Println("------------------- after sorting")
	sort.Ints(l1)
	sort.Ints(l2)
	printLists(l1, l2)

	result, err := similarityScore(l1, l2)
	if err != nil {
		fmt.Println("an error occured!")
		return
	}
	fmt.Println("------------------- similarty score")
	fmt.Println(result)
}
