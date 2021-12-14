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
	out := solution1(example)
	expected := 15
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}

func TestSolution2(t *testing.T) {
	bytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	example := string(bytes)
	out := solution2(example)
	expected := 1134
	if out != expected {
		t.Errorf("want %v, got %v", expected, out)
	}
}
