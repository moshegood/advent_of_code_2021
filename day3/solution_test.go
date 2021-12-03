package main

import (
	"testing"
)

var testInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestSolution(t *testing.T) {
	out := solution(testInput)
	expected := 198
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}

func TestSolution2(t *testing.T) {
	out := solution2(testInput)
	expected := 230
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
