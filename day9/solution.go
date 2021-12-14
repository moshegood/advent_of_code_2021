package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const size = 1000

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	fmt.Println(solution1(input))
	fmt.Println(solution2(input))
}

func inputToHeightMap(input string) [][]int {
	lines := strings.Split(input, "\n")
	heightMap := make([][]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		nums := make([]int, len(line))
		for i, h := range line {
			nums[i] = int(h - '0')
		}
		heightMap = append(heightMap, nums)
	}
	return heightMap
}

func solution1(input string) int {
	heightMap := inputToHeightMap(input)
	mins := []int{}
	for x, row := range heightMap {
		fmt.Println(row)
		for y, h := range row {
			min := true
			if x > 0 && heightMap[x-1][y] <= h {
				min = false
			}
			if x+1 < len(heightMap) && heightMap[x+1][y] <= h {
				min = false
			}
			if y > 0 && row[y-1] <= h {
				min = false
			}
			if y+1 < len(row) && row[y+1] <= h {
				min = false
			}
			if min {
				mins = append(mins, h)
			}
		}
	}
	sum := 0
	for _, m := range mins {
		sum += m + 1
	}
	return sum
}

func solution2(input string) int {
	heightMap := inputToHeightMap(input)
	labels := make([][]int, len(heightMap))
	sames := make(map[int][]int)
	nextLabel := 1
	// Initial pass to label everything.
	// When you see something connected to multiple other groups,
	// log the groups that are the same in the sames variable.
	for i, row := range heightMap {
		labels[i] = make([]int, len(row))
		for j, h := range row {
			connected := []int{}
			if h == 9 {
				continue
			}
			if i > 0 && labels[i-1][j] > 0 {
				connected = append(connected, sames[labels[i-1][j]]...)
			}
			if j > 0 && labels[i][j-1] > 0 {
				connected = append(connected, sames[labels[i][j-1]]...)
			}
			if len(connected) == 0 {
				connected = []int{nextLabel}
				nextLabel++
			}
			sort.Ints(connected)
			labels[i][j] = connected[0]
			uniq := []int{}
			for idx, c := range connected {
				if idx > 0 && c == connected[idx-1] {
					continue
				}
				uniq = append(uniq, c)
			}
			for _, c := range uniq {
				sames[c] = uniq
			}
		}
	}
	// Cleanup to relabel everthing with the smallest label it is attached to.
	// While we are at it, count the number of items with each label.
	fmt.Println("Sames:", sames)
	counts := make(map[int]int)
	for i, row := range labels {
		for j, l := range row {
			if l == 0 {
				continue
			}
			labels[i][j] = sames[l][0]
			counts[labels[i][j]]++
		}
	}
	// Get the product of the sizes of the three biggest groups.
	sizes := []int{}
	for _, v := range counts {
		sizes = append(sizes, v)
	}
	sort.Ints(sizes)
	fmt.Println(sizes)
	lastThree := sizes[len(sizes)-3:]
	fmt.Println(lastThree)
	product := 1
	for _, x := range lastThree {
		product *= x
	}
	return product
}
