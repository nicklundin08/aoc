package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type chain struct {
	checkedSafe []int
	a           int
	b           int
	data        []int
	index       int
	skipsUsed   int
	trend       trend
}

func processNext(c *chain) {
	doesPass := checkTwoElements(c.a, c.b, c.trend)
	if doesPass {
		c.checkedSafe[c.index] = c.a
	}
	c.index++
	c.a = c.data[c.index]
	c.b = c.data[c.index+1]
}

type trend func(a int, b int) bool

var increasing = trend(func(a int, b int) bool {
	return a < b
})

var decreasing = trend(func(a int, b int) bool {
	return a > b
})

func gradual(a int, b int) bool {
	diff := diff(a, b)
	return diff > 0 && diff < 4
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func checkTwoElements(a int, b int, trend trend) bool {
	isGradual := gradual(a, b)
	isTrending := trend(a, b)
	retValue := isGradual && isTrending
	// fmt.Printf("Checking two values %d and %d with trending %t, gradual %t, retval %t.\n", a, b, isTrending, isGradual, retValue)
	return retValue
}

func computeTrend(a []int) trend {
	head := a[0]
	tail := a[1:]
	diff := head - tail[0]
	if diff > 0 {
		return decreasing
	}
	return increasing
}

func checkSliceWithTrend(a []int, trend trend) bool {
	if len(a) == 0 {
		return true
	}
	if len(a) == 1 {
		return true
	}
	head := a[0]
	tail := a[1:]
	doesPass := checkTwoElements(head, tail[0], trend)
	if len(a) == 2 {
		return doesPass
	}
	return doesPass && checkSliceWithTrend(tail, trend)
}

func checkSlice(a []int) bool {
	trend := computeTrend(a)
	return checkSliceWithTrend(a, trend)
}

func check2(data []int) int {
	trend := computeTrend(data)
	for index := 1; index < len(data); index++ {
		a := data[index-1]
		b := data[index]
		doesPass := checkTwoElements(a, b, trend)
		if !doesPass {
			// fmt.Printf("Failure index %d for values %v.\n", index, data)
			return index
		}
	}
	// fmt.Printf("Failure index %d for values %v.\n", -1, data)
	return -1
}

func makeSliceWithoutIndex(d []int, indexToOmit int) []int {
	arr := make([]int, len(d)-1)
	i := 0
	for index, value := range d {
		if indexToOmit != index {
			arr[i] = value
			i++
		}
	}
	return arr
}

func check2WithSkips(data []int) bool {
	if len(data) <= 1 {
		return true
	}
	failureIndex := check2(data)
	fmt.Println(failureIndex)
	if failureIndex == -1 || failureIndex == len(data) {
		return true
	}

	failureIndex2 := check2(makeSliceWithoutIndex(data, failureIndex-1))
	if failureIndex2 == -1 {
		return true
	}

	failureIndex3 := check2(makeSliceWithoutIndex(data, failureIndex))
	if failureIndex3 == -1 {
		return true
	}
	return false
}

func check(head []int, tail []int, skipsRemaining int) bool {
	if skipsRemaining < 0 {
		return false
	}

	if len(tail) == 0 {
		return true
	}
	fmt.Printf("Log head %d tail %v skips %d\n", head, tail, skipsRemaining)
	lastHead := head[len(head)-1]
	diff := lastHead - tail[0]
	trend := increasing
	if diff > 0 {
		trend = decreasing
	}
	if len(tail) == 1 {
		fmt.Println("x")
		return skipsRemaining > 0 || checkTwoElements(lastHead, tail[0], trend)
	}

	doesPass := checkTwoElements(lastHead, tail[0], trend)

	if doesPass {
		fmt.Println("z")
		// return check(tail[0], tail[1:], skipsRemaining)
	}

	skipsRemaining--
	// doesPassIfSkipTail := check(head, tail[1:], skipsRemaining)
	// skipTailAdditionalCheck := check(head, tail[1:1], skipsRemaining)
	// doesPassIfSkipHead := check(tail[0], tail[1:], skipsRemaining)
	// fmt.Printf("skip head: %t. Skip tail %t && %t.\n", doesPassIfSkipHead, doesPassIfSkipTail, skipTailAdditionalCheck)
	// return doesPassIfSkipHead || (doesPassIfSkipTail && skipTailAdditionalCheck)
	return false
}

func rowPasses(row []int, hasDampenerStrategy bool) bool {
	if len(row) <= 2 {
		return false
	}
	skips := 0
	if hasDampenerStrategy {
		skips = 1
	}

	return skips == 7
	// return check(row[0], row[1:], skips)
}
