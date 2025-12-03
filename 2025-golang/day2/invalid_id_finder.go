package day2

import (
	"slices"
	"strconv"
	"strings"

	"github.com/IBM/fp-go/array"
	"github.com/IBM/fp-go/either"
)

type Range struct {
	Lower int
	Upper int
}

func (r *Range) InvalidIds() []int {
	invalidIds := []int{}
	for i := r.Lower; i < r.Upper + 1; i++ {
		if isInvalid(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func (r *Range) InvalidIdSum() int {
	invalids := r.InvalidIds()
	return sum(invalids)
}

func sum(nums []int) int {
	reducer := func(acc, x int) int {
		return acc + x
	}

	return array.Reduce(reducer, 0)(nums)
}

func isInvalid(i int) bool {
	runes := []rune(strconv.Itoa(i))
	return isInvalidForStr(runes)
}

func isInvalidForStr(r []rune) bool {
	if len(r)%2 == 0 {
		return isRepeatedForEvenLength(r)
	}
	return isRepeatedForOddLength(r)
}

func isRepeatedForEvenLength(r []rune) bool {
	halfWayIndex := len(r) / 2
	firstHalf := r[0:halfWayIndex]
	secondHalf := r[halfWayIndex:]
	return slices.Equal(firstHalf, secondHalf)
}

func isRepeatedForOddLength(r []rune) bool {
	return false
}

func BuildRangeFromString(s string) either.Either[error, Range] {
	parts := strings.Split(s, "-")
	lowerStr := parts[0]
	lower, err := strconv.Atoi(lowerStr)
	if err != nil {
		return either.Left[Range](err)
	}
	upperStr := parts[1]
	upper, err := strconv.Atoi(upperStr)
	if err != nil {
		return either.Left[Range](err)
	}
	return either.Right[error](Range{
		Lower: lower,
		Upper: upper,
	})
}
