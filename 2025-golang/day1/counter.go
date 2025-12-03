package day1

import (
	"errors"
	"strconv"
)

type Counter struct {
	Current     int
	Max         int
	ZeroCounter int
	AddOneToZeroCounter bool
}

func MakeCounter(current int) *Counter {
	return &Counter{Current: current, Max: 100, ZeroCounter: 0, AddOneToZeroCounter: true}
}

func (c *Counter) GetZeroCounter() int {
	if c.Current == 0 && c.AddOneToZeroCounter{
		return c.ZeroCounter + 1
	}
	return c.ZeroCounter
}

func (c *Counter) isGTEZero() bool {
	return c.Current >= 0
}

func (c *Counter) isLTMax() bool {
	return c.Current < c.Max
}

func (c *Counter) isValid() bool {
	return c.isGTEZero() && c.isLTMax()
}

// The roll function checks whether the current counter state is valid
// If not it adjusts the current value to be within the bounds and increments the zero counter
// We dont account for landing on zero here - that is done in the `GetZeroCounter` method
// Theres a sneaky edge case where we land right on zero if we are going up (i.e. left) where we then have to ignore the extra landing on zero part
// Not linear time!!!
func (c *Counter) Roll() {
	for !c.isValid() {
		if !c.isGTEZero() {
			c.Current += c.Max
			c.ZeroCounter++
		}
		if !c.isLTMax() {
			c.Current -= c.Max
			c.ZeroCounter++
			// Gross!!
			if c.Current == 0 {
				c.AddOneToZeroCounter = false
			}

		}

	}
}

func (c *Counter) MoveLeft(i int) *Counter {
	c.Current += i
	c.Roll()
	return c
}

func (c *Counter) MoveRight(i int) *Counter {
	c.Current -= i
	c.Roll()
	return c
}

func (c *Counter) Move(s string) (*Counter, error) {
	direction := s[0:1]
	magnitudeString := s[1:]
	magnitude, err := strconv.Atoi(magnitudeString)
	if err != nil {
		return nil, err
	}
	if direction == "L" {
		return c.MoveLeft(magnitude), nil
	}
	if direction == "R" {
		return c.MoveRight(magnitude), nil
	}
	return nil, errors.New("Unknown direction!")
}
