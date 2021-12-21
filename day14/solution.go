package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	template := "FPNFCVSNNFSFHHOCNBOB"
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	fmt.Println(solution(template, input, 10))
	fmt.Println(solution(template, input, 40))
}

type fold struct {
	axis      rune
	foldIndex int
}

const split = " -> "

func parseInput(input string) map[string]string {
	rules := make(map[string]string)
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, split) {
			parts := strings.Split(line, split)
			rules[parts[0]] = parts[1]
		}
	}
	return rules
}

func solution(template string, input string, iterations int) int {
	rules := parseInput(input)

	// Don't forget, we also have the start and end of the template.
	// These are not pairs, but just hang on to the ends.
	pairs := map[string]int{}
	for i := 1; i < len(template); i++ {
		pairs[template[i-1:i+1]]++
	}

	for i := 0; i < iterations; i++ {
		newPairs := make(map[string]int)
		for k, v := range pairs {
			insert := rules[k]
			newPairs[k[:1]+insert] += v
			newPairs[insert+k[1:]] += v
		}
		pairs = newPairs
	}

	// Count the letters
	counts := make(map[byte]int)
	for k, v := range pairs {
		counts[k[0]] += v
		counts[k[1]] += v
	}
	// Each letter was double counted because it is in two pairs
	for k, v := range counts {
		counts[k] = v / 2
	}
	// Count the ends once each
	counts[template[0]]++
	counts[template[len(template)-1]]++

	// Sort and compare max to min
	values := []int{}
	for _, v := range counts {
		values = append(values, v)
	}
	sort.Ints(values)
	return values[len(values)-1] - values[0]
}
