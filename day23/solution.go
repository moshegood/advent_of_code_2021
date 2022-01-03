package main

import (
	"fmt"
	"math"
)

var allDone [11]string

var froms []int = []int{2, 4, 6, 8}
var tos []int = []int{0, 1, 3, 5, 7, 9, 10}
var costs map[string]int = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

var finalSpots map[string]int = map[string]int{
	"A": 2,
	"B": 4,
	"C": 6,
	"D": 8,
}

func main() {
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
	cost := solution(part1Input, nil) + baseCost(part1Input)
	fmt.Println("Total cost:", cost)

	part2Input := [11]string{
		"",
		"",
		"BDDB",
		"",
		"CCBC",
		"",
		"ABAD",
		"",
		"DACA",
		"",
		"",
	}
	cost2 := solution(part2Input, nil) + baseCost(part2Input)
	fmt.Println("Part2 total cost:", cost2)
}

func baseCost(input [11]string) int {
	cost := 0
	maxL := 0
	for _, s := range input[:] {
		if len(s) > maxL {
			maxL = len(s)
		}
		for i, c := range s {
			cost += int(math.Pow(10, float64(c-'A'))) * (i + 1)
		}
	}
	cost += 1111 * maxL * (maxL + 1) / 2
	return cost
}

// pathExists checks to see if you can move from "from" to "to".
// But does not check the end points themselves.
func pathExists(input [11]string, from, to int) bool {
	toMap := map[int]bool{
		0:  true,
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		9:  true,
		10: true,
	}

	direction := 1
	if from < to {
		direction *= -1
	}
	for i := to + direction; i != from; i += direction {
		if toMap[i] && input[i] != "" {
			return false
		}
	}

	return true
}

func solution(input [11]string, seen map[[11]string]int) int {
	if seen == nil {
		seen = make(map[[11]string]int)
	}
	if v, ok := seen[input]; ok {
		return v
	}

	// Move anything to its final resting place if at all possible.
	// Always do this if we possibly can.
	for _, to := range tos {
		if input[to] == "" {
			continue
		}
		dest := finalSpots[input[to]]
		if input[dest] == "" && pathExists(input, dest, to) {
			letterMoved := input[to]
			input[to] = ""
			return costs[letterMoved]*abs(dest-to) + solution(input, seen)
		}
	}

	if input == allDone {
		return 0
	}
	bestSolution := int(1e9)
	for _, from := range froms {
		if input[from] == "" {
			continue
		}

		for _, to := range tos {
			if input[to] != "" {
				continue
			}

			if !pathExists(input, from, to) {
				continue
			}

			// Try it recusively
			testInput := input
			testInput[to] = testInput[from][:1]
			testInput[from] = testInput[from][1:]
			// fmt.Println("Moving from:", from, "to:", to)
			possible := solution(testInput, seen)
			possible += costs[testInput[to]] * abs(from-to)
			if possible < bestSolution {
				bestSolution = possible
			}
		}
	}

	seen[input] = bestSolution
	return bestSolution
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
