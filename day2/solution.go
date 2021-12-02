package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileInput := readInput()
	fmt.Println(solution(fileInput))
	fmt.Println(solution2(fileInput))
}

type row struct {
	direction string
	magnitude int
}

func readInput() []row {
	fileHandle, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	var values []row
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		values = append(values, row{parts[0], value})
	}
	return values
}

func solution2(values []row) int {
	aim := 0
	depth := 0
	dist := 0

	for _, value := range values {
		switch value.direction {
		case "down":
			aim += value.magnitude
		case "up":
			aim -= value.magnitude
		case "forward":
			dist += value.magnitude
			depth += value.magnitude * aim
		}
	}
	return int(math.Abs(float64(depth) * float64(dist)))
}

func solution(values []row) int {
	depth := 0
	dist := 0

	for _, value := range values {
		switch value.direction {
		case "down":
			depth += value.magnitude
		case "up":
			depth -= value.magnitude
		case "forward":
			dist += value.magnitude
		}
	}
	return int(math.Abs(float64(depth) * float64(dist)))
}
