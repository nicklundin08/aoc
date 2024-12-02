package main

type rule func(a int, b int) bool

var increasing = rule(func(a int, b int) bool {
	return a < b
})

var decreasing = rule(func(a int, b int) bool {
	return a > b
})

var gradual = rule(func(a int, b int) bool {
	diff := diff(a, b)
	return diff > 0 && diff < 4
})

func diff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func rowPasses(row []int) bool {
	if len(row) == 0 {
		return false
	}
	diff := row[0] - row[len(row)-1]
	trendRule := increasing
	if diff > 0 {
		trendRule = decreasing
	}
	for i := 0; i < len(row)-1; i++ {
		a := row[i]
		b := row[i+1]
		passes := passes(a, b, trendRule)
		if !passes {
			return false
		}
	}
	return true
}

func passes(a int, b int, trendRule rule) bool {
	isGradual := gradual(a, b)
	if !isGradual {
		return false
	}
	isConsistent := trendRule(a, b)
	if !isConsistent {
		return false
	}
	return isGradual && isConsistent
}
