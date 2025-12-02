package day1

import (
	"errors"
	"strconv"
)

type Counter struct {
	Current     int
	Max         int
	ZeroCounter int
}

func MakeCounter(current int) Counter {
	return Counter{Current: current, Max: 100, ZeroCounter: 0}
}

func (c Counter) NewCounter(current int) Counter {
	if current < 0 {
		return c.NewCounter(current + c.Max)
	}
	if current >= c.Max {
		return c.NewCounter(current - c.Max)
	}
	if current == 0 {
		c.ZeroCounter++
	}
	return Counter{Current: current, Max: c.Max, ZeroCounter: c.ZeroCounter}
}

func (c Counter) MoveLeft(i int) Counter {
	return c.NewCounter(c.Current + i)
}

func (c Counter) MoveRight(i int) Counter {
	return c.NewCounter(c.Current - i)
}

func (c Counter) Move(s string) (Counter, error) {
	direction := s[0:1]
	magnitudeString := s[1:]
	magnitude, err := strconv.Atoi(magnitudeString)
	if err != nil {
		return Counter{}, err
	}
	if direction == "L" {
		return c.MoveLeft(magnitude), nil
	}
	if direction == "R" {
		return c.MoveRight(magnitude), nil
	}
	return Counter{}, errors.New("Unknown direction!");
}
