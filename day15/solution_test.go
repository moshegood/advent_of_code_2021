package main

import (
	"os"
	"testing"
)

func TestSolution(t *testing.T) {

	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	result := solution(example, 1)
	expected := 40
	if result != expected {
		t.Errorf("want %v, got %v", expected, result)
	}
}

func TestSolutionPart2(t *testing.T) {

	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	result := solution(example, 5)
	expected := 315
	if result != expected {
		t.Errorf("want %v, got %v", expected, result)
	}
}
