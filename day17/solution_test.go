package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	result := part1(-10)
	expected := 45
	if result != expected {
		t.Errorf("want %v, got %v", expected, result)
	}
}

func TestSolutionPart2(t *testing.T) {

	xmin := 20
	xmax := 30
	ymin := -10
	ymax := -5

	expected := 112
	result := part2(xmin, xmax, ymin, ymax)
	if result != expected {
		t.Errorf("want %v, got %v", expected, result)
	}

}
