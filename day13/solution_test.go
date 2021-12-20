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
	sum := len(solution(example, 1))
	expectedSum := 17
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}

	solution(example, 2)
}
