package main

type safetyCaclulator struct {
	// Either increasing or decreasing
	trend trend
	// The row to inspect
	row []int
	// IndexA walks the array as the lagging index
	indexA int
	// IndexB walks the array as the leading index
	indexB int
	//Dampening is implemented as a lives remaning strategy. Normally during the walk we would increment both indexes, however if an pair of eleemnts doesnt pass, we "skip" the offending element by incrementing IndexB only
	//[1,2,7,3,5]
	// a b
	//[1,2,7,3,5]
	//   a b
	//[1,2,7,3,5]
	//   a   b        <---- b increemnts 1, a is not changed
	//[1,2,7,3,5]
	//       a b      <---- b increments 1, a is reset to b -1
	livesRemaning int
}

func (sc *safetyCaclulator) check() {
	if sc.indexA == len(sc.row) || sc.indexB == len(sc.row) {
		return
	}
	if sc.livesRemaning == 0 {
		return
	}
	a := sc.row[sc.indexA]
	b := sc.row[sc.indexB]
	doesPass := checkTwoElements(a, b, sc.trend)
	sc.indexB++
	if doesPass {
		sc.indexA = sc.indexB - 1
	} else {
		sc.livesRemaning--
	}
	sc.check()
}

func createNewSafetyCalculator(row []int, withDampener bool, trend trend) *safetyCaclulator {
	lives := 1

	if withDampener {
		lives++
	}

	return &safetyCaclulator{
		row:           row,
		trend:         trend,
		indexA:        0,
		indexB:        1,
		livesRemaning: lives,
	}
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
	// fmt.Printf("Logging %d, %d.\n", a, b)
	isGradual := gradual(a, b)
	if !isGradual {
		return false
	}
	isTrending := trend(a, b)
	return isTrending
}

func rowPasses(row []int, hasDampenerStrategy bool) bool {
	if len(row) <= 2 {
		return false
	}
	diff := row[0] - row[len(row)-1]
	trend := increasing
	if diff > 0 {
		trend = decreasing
	}
	sc := createNewSafetyCalculator(row, hasDampenerStrategy, trend)
	sc.check()
	return sc.livesRemaning > 0
}
