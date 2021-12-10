package main

import (
	"fmt"
	"sort"
)

const size = 1000

func main() {
	fmt.Println(solution1(input))
	fmt.Println(solution2(input))
}

func solution1(input []int) int {
	sort.Ints(input)
	medianLowSide := input[len(input)/2]
	medianHighSide := input[len(input)/2+1]
	fmt.Println("median:", medianLowSide, "->", medianHighSide)
	bestOptions := []int{medianLowSide, medianHighSide}

	return bestOption(input, bestOptions, func(x int) int { return x })
}

func bestOption(input []int, destinations []int, costFunc func(int) int) int {
	minCost := int(1e16)
	for _, dest := range destinations {
		cost := totalCost(input, dest, costFunc)
		if cost < minCost {
			minCost = cost
		}
		fmt.Println("dest:", dest, "cost:", cost)
	}
	return minCost
}

func totalCost(input []int, destination int, costFunc func(int) int) int {
	cost := 0
	for _, i := range input {
		distance := 0
		if i > destination {
			distance = i - destination
		}
		if i < destination {
			distance = destination - i
		}
		cost += costFunc(distance)
	}
	return cost
}

func solution2(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	floorAvg := sum / len(input)
	bestOptions := []int{floorAvg, floorAvg + 1}
	fmt.Println("avg:", floorAvg, "sum:", sum, "len:", len(input))
	fmt.Println("best:", bestOptions)

	return bestOption(input, bestOptions, func(x int) int { return x * (x + 1) / 2 })
}
