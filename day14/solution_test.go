package main

import (
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	template := "NNCB"

	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	result := solution(template, example, 10)
	expected := 1588
	if result != expected {
		t.Errorf("want %v, got %v", expected, result)
	}
}

func TestSolutionPart2(t *testing.T) {
	template := "NNCB"

	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	result := solution(template, example, 40)
	expected := 2188189693529
	if result != expected {
		t.Errorf("want %v, got %v", expected, result)
	}
}
