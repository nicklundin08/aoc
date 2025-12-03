package day1

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test__PrintAnswer(t *testing.T) {
	answer, err := getAnswer("input.txt")
	if err != nil {
		fmt.Println("An error occured", err)
	}

	fmt.Println(answer)
}

func Test__SuppliedCaseWithFile(t *testing.T) {
	answer, err := getAnswer("sample.txt")
	if err != nil {
		fmt.Println("An error occured", err)
	}

	want := 6
	if answer != want {
		t.Errorf(`Default Case = %d, want %d`, answer, want)
	}
}

func Test__SuppliedCase(t *testing.T) {
	answer := MakeCounter(50).
		MoveLeft(68).
		MoveLeft(30).
		MoveRight(48).
		MoveLeft(5).
		MoveRight(60).
		MoveLeft(55).
		MoveLeft(1).
		MoveLeft(99).
		MoveRight(14).
		MoveLeft(82).
		GetZeroCounter()

	want := 6
	if answer != want {
		t.Errorf(`Default Case = %d, want %d`, answer, want)
	}
}

func TestIsValid(t *testing.T) {
	if !MakeCounter(50).isValid() {
		t.Errorf(`Counter(50) should be valid`)
	}

	if MakeCounter(-1).isValid() {
		t.Errorf(`Counter(-1) should not be valid`)
	}

	if MakeCounter(100).isValid() {
		t.Errorf(`Counter(100) should not be valid`)
	}
	if !MakeCounter(99).isValid() {
		t.Errorf(`Counter(99) should be valid`)
	}
	if !MakeCounter(0).isValid() {
		t.Errorf(`Counter(0) should be valid`)
	}
}

func TestMoveRight__DontRoll(t *testing.T) {
	c := MakeCounter(50).MoveRight(1)

	if c.Current != 49 {
		t.Errorf(`Wanted Counter.Current to be 49 but was %d`, c.Current)
	}

	if c.GetZeroCounter() != 0 {
		t.Errorf(`Wanted Counter.GetZeroCounter() to be 0 but was %d`, c.GetZeroCounter())
	}
}

func TestMoveLeft__DontRoll(t *testing.T) {
	c := MakeCounter(50).MoveLeft(1)

	if c.Current != 51 {
		t.Errorf(`Wanted Counter.Current to be 51 but was %d`, c.Current)
	}

	if c.GetZeroCounter() != 0 {
		t.Errorf(`Wanted Counter.GetZeroCounter() to be 0 but was %d`, c.GetZeroCounter())
	}
}

func TestMoveRight__Roll(t *testing.T) {
	c := MakeCounter(50).
		MoveRight(50).
		MoveRight(100)

	if c.Current != 0 {
		t.Errorf(`Wanted Counter.Current to be 0 but was %d`, c.Current)
	}

	if c.GetZeroCounter() != 2 {
		t.Errorf(`Wanted Counter.GetZeroCounter() to be 2 but was %d`, c.GetZeroCounter())
	}
}

func TestMoveLeft__Roll(t *testing.T) {
	c := MakeCounter(50).
		MoveLeft(50).
		MoveLeft(100)

	if c.Current != 0 {
		t.Errorf(`Wanted Counter.Current to be 0 but was %d`, c.Current)
	}

	if c.GetZeroCounter() != 2 {
		t.Errorf(`Wanted Counter.GetZeroCounter() to be 2 but was %d`, c.GetZeroCounter())
	}
}

func getAnswer(path string) (int, error) {
	counter := MakeCounter(50)
	moves, err := readFile(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, move := range moves {
		counter, err = counter.Move(move)
		if err != nil {
			return 0, err
		}
	}
	return counter.GetZeroCounter(), nil
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}
