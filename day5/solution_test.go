package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	out := solution(example, false)
	expected := 5
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
func TestSolutionPart2(t *testing.T) {
	out := solution(example, true)
	expected := 12
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
