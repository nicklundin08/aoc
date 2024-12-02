package main

import (
	"testing"
)

func TestThingy(t *testing.T) {
	result := 2 + 4
	expected := 6

	if result != expected {
		t.Errorf("wanted %d but found %d", expected, result)
	}
}

func TestThingy2(t *testing.T) {
	result := 2 + 4
	expected := 6

	if result != expected {
		t.Errorf("wanted %d but found %d", expected, result)
	}
}
