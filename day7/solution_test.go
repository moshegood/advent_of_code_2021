package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	out := solution1(example)
	expected := 37
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
func TestSolutionPart2(t *testing.T) {
	out := solution2(example)
	expected := 168
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
