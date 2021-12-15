package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	fmt.Println(solution(input))
}

func solution(input string) (int, int) {
	matches := map[rune]rune{
		'>': '<',
		')': '(',
		']': '[',
		'}': '{',
	}
	points := map[rune]int{
		'>': 25137,
		')': 3,
		']': 57,
		'}': 1197,
	}
	scores := map[rune]int{
		'<': 4,
		'(': 1,
		'[': 2,
		'{': 3,
	}
	lines := strings.Split(input, "\n")
	sum := 0
	incompleteScores := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		isBroken := false
		currentOpens := []rune{}
		for _, c := range line {
			if strings.ContainsRune("<({[", c) {
				currentOpens = append(currentOpens, c)
			}
			if strings.ContainsRune(">]})", c) {
				if currentOpens[len(currentOpens)-1] == matches[c] {
					currentOpens = currentOpens[:len(currentOpens)-1]
				} else {
					// broken!!
					sum += points[c]
					// fmt.Println("Broke:", line[:i], string(c))
					isBroken = true
					break
				}
			}
		}
		if !isBroken {
			// fmt.Println("Current opens:", string(currentOpens))
			score := 0
			for len(currentOpens) > 0 {
				last := len(currentOpens) - 1
				score *= 5
				score += scores[currentOpens[last]]
				currentOpens = currentOpens[:last]
			}
			incompleteScores = append(incompleteScores, score)
		}
	}
	sort.Ints(incompleteScores)
	fmt.Println("Scores:", incompleteScores)
	middle := incompleteScores[len(incompleteScores)/2]
	fmt.Println("Middle:", middle)
	return sum, middle
}
