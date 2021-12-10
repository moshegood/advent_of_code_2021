package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	out := solution(example, 80)
	expected := 5934
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
func TestSolutionPart2(t *testing.T) {
	out := solution(example, 256)
	expected := 26984457539
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
