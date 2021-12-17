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
	sum := solution(example, 0)
	expectedSum := 10
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}

	iterations := solution(example, 1)
	expectedIter := 36
	if iterations != expectedIter {
		t.Errorf("want %v, got %v", expectedIter, iterations)
	}
}
