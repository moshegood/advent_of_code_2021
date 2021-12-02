package main

import (
	"testing"
)

var testInput = []row{
	{"forward", 5},
	{"down", 5},
	{"forward", 8},
	{"up", 3},
	{"down", 8},
	{"forward", 2},
}

func TestSolution(t *testing.T) {
	out := solution(testInput)
	expected := 150
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
	out = solution2(testInput)
	expected = 900
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
