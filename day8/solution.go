package main

import (
	"fmt"
	"strings"
)

const size = 1000

func main() {
	fmt.Println(solution1(input))
	fmt.Println(solution2(input))
}

func solution1(input []string) int {
	count := 0
	for _, line := range input {
		words := strings.Fields(line)
		// theTen := words[:10]
		output := words[11:]
		for _, letter := range output {
			switch len(letter) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func solution2(input []string) int {
	sum := 0
	for _, line := range input {
		words := strings.Fields(line)
		theTen := words[:10]
		decoder := decode(theTen)

		// fmt.Printf("Decoder: %+v\n", decoder)
		output := words[11:]
		value := 0
		for _, word := range output {
			value *= 10
			value += decoder[toBitmask(word)]
			// fmt.Printf("%q => %v\n", word, decoder[toBitmask(word)])
		}
		sum += value
		// fmt.Println("Value:", value)
	}
	return sum
}

func toBitmask(input string) (out uint8) {
	for _, c := range input {
		out += 1 << (c - 'a')
	}
	return
}

func onesInMask(input uint8) int {
	count := 0
	for input != 0 {
		count++
		input &= (input - 1)
	}
	return count
}

func decode(theTen []string) map[uint8]int {
	intToMask := make(map[int]uint8)
	output := make(map[uint8]int)
	byLen := [8][]string{}
	for _, word := range theTen {
		byLen[len(word)] = append(byLen[len(word)], word)
	}

	for length, list := range byLen {
		for _, word := range list {
			mask := toBitmask(word)
			switch length {
			case 2:
				intToMask[1] = mask
				output[mask] = 1
			case 3:
				output[mask] = 7
			case 4:
				intToMask[4] = mask
				output[mask] = 4
			case 5:
				if intToMask[1]&mask == intToMask[1] {
					output[mask] = 3
				} else if onesInMask(intToMask[4]&mask) == 2 {
					output[mask] = 2
				} else {
					output[mask] = 5
				}
			case 6:
				if intToMask[4]&mask == intToMask[4] {
					output[mask] = 9
				} else if intToMask[1]&mask == intToMask[1] {
					output[mask] = 0
				} else {
					output[mask] = 6
				}
			case 7:
				output[mask] = 8
			}
		}
	}
	return output
}
