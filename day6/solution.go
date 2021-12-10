package main

import (
	"fmt"
)

const size = 1000

func main() {
	fmt.Println(solution(input, 80))
	fmt.Println(solution(input, 256))
}

func solution(input []int, days int) int {
	countsPerDay := []int{8: 0}
	for _, value := range input {
		countsPerDay[value]++
	}
	for i := 0; i < days; i++ {
		zeros := countsPerDay[0]
		countsPerDay = countsPerDay[1:]
		countsPerDay[6] += zeros
		countsPerDay = append(countsPerDay, zeros)
	}
	sum := 0
	for _, count := range countsPerDay {
		sum += count
	}
	return sum
}
