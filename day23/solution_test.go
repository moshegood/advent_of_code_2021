package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	part1Input := [11]string{
		"",
		"",
		"BB",
		"",
		"CC",
		"",
		"AD",
		"",
		"DA",
		"",
		"",
	}
	value := solution(part1Input, nil) + baseCost(part1Input)
	expected := 11120
	if value != expected {
		t.Errorf("solution(...) = %v, want %v", value, expected)
	}
}
