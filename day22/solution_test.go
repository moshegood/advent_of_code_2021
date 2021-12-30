package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	sum := part1(example, 50)
	expectedSum := 590784
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}
}

func TestPart2(t *testing.T) {
	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	sum := part2(example)
	expectedSum := 590784
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}
}

func TestPart2BigInput(t *testing.T) {
	bytes, err := os.ReadFile("example2.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	sum := part2(example)
	expectedSum := 2758514936282235
	if sum != expectedSum {
		t.Errorf("want %v, got %v", expectedSum, sum)
	}
}
