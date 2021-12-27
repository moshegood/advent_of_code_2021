package main

import "fmt"

func main() {
	xmin := 175
	xmax := 227
	ymin := -134
	ymax := -79

	fmt.Println("xmin:", xmin)
	fmt.Println("xmax:", xmax)
	fmt.Println("ymin:", ymin)
	fmt.Println("ymax:", ymax)
	fmt.Println("Part1:", part1(ymin))
	fmt.Println("Part2:", part2(xmin, xmax, ymin, ymax))
}

func part1(ymin int) int {
	return ymin * (ymin + 1) / 2
}

func part2(xmin, xmax, ymin, ymax int) int {
	maxSteps := 2 * (-ymin)
	maxY0 := (-ymin)
	minY0 := ymin

	xByStep := map[int]map[int]bool{}
	for x0 := 0; x0 <= xmax; x0++ {
		x := 0
		for step := 0; step <= maxSteps; step++ {
			if xByStep[step] == nil {
				xByStep[step] = make(map[int]bool)
			}
			if step < x0 {
				x += x0 - step
			}
			if x >= xmin && x <= xmax {
				xByStep[step][x0] = true
			}
		}
	}

	possibilites := 0
	for y0 := minY0; y0 <= maxY0; y0++ {
		y := 0
		xs := map[int]bool{}
		for step := 0; y >= ymin; step++ {
			y += y0 - step
			if y >= ymin && y <= ymax {
				for x := range xByStep[step] {
					xs[x] = true
				}
			}
		}
		for _ = range xs {
			possibilites++
		}
	}

	return possibilites
}
