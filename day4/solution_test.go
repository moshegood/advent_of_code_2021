package main

import (
	"testing"
)

var testInputFile = "example.txt"

func TestSolution(t *testing.T) {
	input := readInput(testInputFile)
	out := solution(input)
	expected := 4512
	if out[0] != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
