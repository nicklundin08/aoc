package main

import (
	"errors"
	"fmt"
	"sort"
  "os"
  "strings"
  "strconv"
)

func abs(x int) int {
	if x > 0 {
		return x
	}

	return 0 - x
}

func readListsFromFile(path string)([]int, []int, error){
  fileBytes, err := os.ReadFile(path)
  if err != nil{
    return nil, nil, err
  }
  fileStr := string(fileBytes)
  values := strings.Fields(fileStr) 

  l1 := make([]int, 0)
  l2 := make([]int, 0)
  for index, value := range values{
    intVal, err := strconv.Atoi(value)
    if err != nil {
      return nil, nil, err
    }
    if index % 2 == 0{
      l1 = append(l1, intVal)
    }else{
      l2 = append(l2, intVal)
    }
  }
  return l1, l2, nil
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
  path := "./input.txt"
  l1, l2, err := readListsFromFile(path)
  if err != nil{
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
