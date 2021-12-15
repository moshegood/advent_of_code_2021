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
	sum, middle := solution(example)
	expectedSum := 26397
	expectedMid := 288957
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}
	if middle != expectedMid {
		t.Errorf("want %v, got %v", expectedMid, middle)
	}
}
