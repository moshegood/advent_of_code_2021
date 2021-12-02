package main

import (
	"testing"
)

var testInput = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestSolution(t *testing.T) {
	out := solution(testInput, 1)
	if out != 7 {
		t.Errorf("No window error. Expected 7, got %v", out)
	}
	out = solution(testInput, 3)
	if out != 5 {
		t.Errorf("Window error. Expected 5, got %v", out)
	}
}
