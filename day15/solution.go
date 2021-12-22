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
	fmt.Println(solution(input, 1))
	fmt.Println(solution(input, 5))
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

func blowupGrid(grid [][]int, times int) [][]int {
	newGrid := make([][]int, len(grid)*times)
	for s := 0; s < times; s++ {
		for i, row := range grid {
			newRow := make([]int, len(row)*times)
			for t := 0; t < times; t++ {
				for j, value := range row {
					newRow[t*len(row)+j] = (value+t+s-1)%9 + 1
				}
			}
			newGrid[s*len(grid)+i] = newRow
		}
	}
	return newGrid
}

func solution(input string, iterations int) int {
	grid := parseInput(input)
	grid = blowupGrid(grid, iterations)

	bigValue := 100 * len(grid) * len(grid[0])

	// Initialize to huge costs everywhere
	costs := make([][]int, len(grid))
	for i := range costs {
		costs[i] = make([]int, len(grid[0]))
		for j := range costs[i] {
			costs[i][j] = bigValue
		}
	}
	// Except the initial spot, which you start in for free
	costs[0][0] = 0

	// Iterate back and forth until we stabilize
	changed := true
	reversedX := false
	reversedY := false
	for changed {
		changed = false
		for i, row := range grid {
			for j := range row {
				if reversedX {
					i = len(grid) - 1 - 1
				}
				if reversedY {
					j = len(row) - 1 - j
				}
				p := point{i, j}
				neighbors := []int{}
				for _, q := range neighborsOf(grid, p) {
					neighbors = append(neighbors, costs[q.x][q.y])
				}
				sort.Ints(neighbors)
				min := neighbors[0]
				if p.x == 0 && p.y == 0 {
					min = 0
				}
				existingCost := costs[p.x][p.y]
				costs[p.x][p.y] = min + grid[p.x][p.y]
				if existingCost != costs[p.x][p.y] {
					changed = true
				}
			}
		}

		// Iterate through the different directions. Rotate 90 each time.
		// This allows quick propagation of changes in different directions.
		if !reversedX && !reversedY {
			reversedX = true
		} else if reversedX && !reversedY {
			reversedY = true
		} else if reversedX && reversedY {
			reversedX = false
		} else {
			reversedY = false
		}
	}
	return costs[len(costs)-1][len(costs[0])-1] - costs[0][0]
}

type point struct {
	x, y int
}

func neighborsOf(grid [][]int, p point) []point {
	ps := []point{}
	if p.x > 0 {
		ps = append(ps, point{p.x - 1, p.y})
	}
	if p.y > 0 {
		ps = append(ps, point{p.x, p.y - 1})
	}
	if p.x+1 < len(grid) {
		ps = append(ps, point{p.x + 1, p.y})
	}
	if p.y+1 < len(grid[0]) {
		ps = append(ps, point{p.x, p.y + 1})
	}
	return ps
}
