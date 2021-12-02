package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileInput := readInput()
	fmt.Println(solution(fileInput, 1))
	fmt.Println(solution(fileInput, 3))
}

func readInput() []int {
	fileHandle, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	var values []int
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}
	return values
}

func solution(values []int, window int) int {
	c := 0
	for i := window; i < len(values); i++ {
		if values[i] > values[i-window] {
			c++
		}
	}
	return c
}
