package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	fmt.Println(len(solution(input, 1)))
	fmt.Println(len(solution(input, 100)))
}

type fold struct {
	axis      rune
	foldIndex int
}

const foldPrefix = "fold along "

func parseInput(input string) (dots [][2]int, folds []fold) {
	for _, line := range strings.Split(input, "\n") {
		if strings.ContainsRune(line, ',') {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			dots = append(dots, [2]int{x, y})
		} else if strings.HasPrefix(line, foldPrefix) {
			line = line[len(foldPrefix):]
			parts := strings.Split(line, "=")
			v, _ := strconv.Atoi(parts[1])
			folds = append(folds, fold{rune(parts[0][0]), v})
		}
	}
	return
}

func solution(input string, max_folds int) [][2]int {
	dots, folds := parseInput(input)

	dotMap := make(map[[2]int]bool)
	for _, d := range dots {
		dotMap[d] = true
	}

	for i, f := range folds {
		if i >= max_folds {
			break
		}
		foldedDots := map[[2]int]bool{}
		for _, d := range dots {
			index := int(f.axis - 'x')
			if d[index] > f.foldIndex {
				d[index] = 2*f.foldIndex - d[index]
			}
			foldedDots[d] = true
		}
		dots = dots[:0]
		for k := range foldedDots {
			dots = append(dots, k)
		}
	}

	printDots(dots)
	return dots
}

func printDots(dots [][2]int) {
	var max_x, max_y int
	dotMap := make(map[int]map[int]bool)
	for _, d := range dots {
		if x := d[0]; x > max_x {
			max_x = x
		}
		if y := d[1]; y > max_y {
			max_y = y
		}
		if _, ok := dotMap[d[0]]; !ok {
			dotMap[d[0]] = make(map[int]bool)
		}
		dotMap[d[0]][d[1]] = true
	}
	for y := 0; y <= max_y; y++ {
		for x := 0; x <= max_x; x++ {
			c := " "
			if dotMap[x][y] {
				c = "█"
			}
			fmt.Print(c)
		}
		fmt.Println("")
	}
}
