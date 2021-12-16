package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	fmt.Println(solution(input, 100))
	fmt.Println(solution(input, 300))
}

func parseInput(input string) [][]int {
	rows := [][]int{}
	for _, line := range strings.Fields(input) {
		row := []int{}
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		rows = append(rows, row)
	}
	return rows
}

func solution(input string, maxIterations int) (int, int) {
	rows := parseInput(input)
	total := 0
	i := 0
	for i = 0; i < maxIterations; i++ {
		// Step 1. Increase add octopi
		for x := range rows {
			// fmt.Printf("%+v\n", rows[x])
			for y := range rows[x] {
				rows[x][y]++
			}
		}
		// Step 2. Flash and increase neighbors
		flashed := make([][]bool, len(rows))
		for x := range rows {
			flashed[x] = make([]bool, len(rows[0]))
		}

		for {
			anyFlashes := false
			for x := range rows {
				for y := range rows[x] {
					if !flashed[x][y] && rows[x][y] > 9 {
						flashed[x][y] = true
						anyFlashes = true
						total++
						for dx := -1; dx <= 1; dx++ {
							if x+dx < 0 || x+dx >= len(rows) {
								continue
							}
							for dy := -1; dy <= 1; dy++ {
								if y+dy < 0 || y+dy >= len(rows[0]) {
									continue
								}
								rows[x+dx][y+dy]++
							}
						}
					}
				}
			}

			if !anyFlashes {
				break
			}
		}
		// Step 3. Reset flashed octopi
		allZero := true
		for x := range rows {
			for y := range rows[x] {
				if rows[x][y] > 9 {
					rows[x][y] = 0
				} else {
					allZero = false
				}
			}
		}

		// Abort on all zeros
		if allZero {
			// But add one first, because this iteration was complete
			i++
			break
		}
	}
	return total, i
}
