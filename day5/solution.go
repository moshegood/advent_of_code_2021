package main

import (
	"fmt"
)

const size = 1000

func main() {
	fmt.Println(solution(input, false))
	fmt.Println(solution(input, true))
}

type fileData struct {
	lines [][2][2]int
}

func solution(input fileData, withDiagonals bool) int {
	grid := [size][size]int{}
	for _, line := range input.lines {
		x1, y1, x2, y2 := line[0][0], line[0][1], line[1][0], line[1][1]
		var delta_x, delta_y int
		if x1 > x2 {
			delta_x = -1
		} else if x2 > x1 {
			delta_x = 1
		}
		if y1 > y2 {
			delta_y = -1
		} else if y2 > y1 {
			delta_y = 1
		}
		x, y := x1, y1
		if !withDiagonals {
			if x != x2 && y != y2 {
				continue
			}
		}
		for {
			grid[x][y]++
			if x == x2 && y == y2 {
				break
			}
			x += delta_x
			y += delta_y
		}
	}
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] > 1 {
				sum++
			}
		}
	}
	return sum
}
