package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileInput := readInput()
	fmt.Println(solution(fileInput))
	fmt.Println(solution2(fileInput))
}

func readInput() []string {
	fileHandle, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	var values []string
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values = append(values, line)
	}
	return values
}

func mostAndLeastCommonBits(values []string) (string, string) {
	size := len(values[0])
	counts := [][2]int{}
	for i := 0; i < size; i++ {
		counts = append(counts, [2]int{})
	}
	for _, value := range values {
		for i, c := range value {
			counts[i][c-'0']++
		}
	}
	fmt.Println(counts)
	mins := ""
	maxs := ""
	for _, c := range counts {
		if c[0] > c[1] {
			mins += "1"
			maxs += "0"
		} else {
			mins += "0"
			maxs += "1"
		}
	}
	return maxs, mins
}

func solution(values []string) int {
	maxs, mins := mostAndLeastCommonBits(values)
	min, err := strconv.ParseInt(mins, 2, 64)
	if err != nil {
		panic(err)
	}
	max, err := strconv.ParseInt(maxs, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(min * max)
}

func findRecursivePrefixThing(values []string, compare func(int, int) byte) string {
	kept := values
	for i := 0; i < len(values[0]); i++ {
		counts := [2]int{}
		for _, value := range kept {
			counts[value[i]-'0']++
		}
		valueToKeep := compare(counts[0], counts[1])
		keep := []string{}
		for _, value := range kept {
			if value[i] == valueToKeep {
				keep = append(keep, value)
			}
		}
		kept = keep
		if len(kept) == 1 {
			return kept[0]
		}
	}
	return ""
}

func solution2(values []string) int {
	this := findRecursivePrefixThing(values, func(x, y int) byte {
		moreCommon := '1'
		if x > y {
			moreCommon = '0'
		}
		return byte(moreCommon)
	})
	that := findRecursivePrefixThing(values, func(x, y int) byte {
		lessCommon := '0'
		if x > y {
			lessCommon = '1'
		}
		return byte(lessCommon)
	})

	fmt.Printf("%q, %q\n", this, that)

	x, err := strconv.ParseInt(this, 2, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(that, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(x * y)
}
