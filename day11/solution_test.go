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
	sum, _ := solution(example, 100)
	expectedSum := 1656
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}
	_, iterations := solution(example, 200)
	expectedIter := 195
	if iterations != expectedIter {
		t.Errorf("want %v, got %v", expectedIter, iterations)
	}
}
