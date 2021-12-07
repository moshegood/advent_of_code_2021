package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

func main() {
	fileInput := readInput("input.txt")
	fmt.Println(solution(fileInput))
}

type board struct {
	numbers [boardSize][boardSize]int
	wins    [2 * boardSize][boardSize]int
}

type fileData struct {
	picks  []int
	boards []board
}

func readInput(filename string) fileData {
	fileHandle, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	data := fileData{}
	nextBoard := board{}
	row := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if data.picks == nil {
			numbers := strings.Split(line, ",")
			for _, n := range numbers {
				v, _ := strconv.Atoi(n)
				data.picks = append(data.picks, v)
			}
			continue
		}
		numbers := strings.Fields(line)
		if len(numbers) == 0 {
			continue
		}
		for i, n := range numbers {
			v, _ := strconv.Atoi(n)
			nextBoard.numbers[row][i] = v
		}
		row++
		if row == boardSize {
			row = 0
			for i := 0; i < boardSize; i++ {
				// The five across
				nextBoard.wins[i] = nextBoard.numbers[i]
				// The columns
				for j := 0; j < boardSize; j++ {
					nextBoard.wins[boardSize+i][j] = nextBoard.numbers[j][i]
				}
				// Diagonals
				// nextBoard.wins[10][i] = nextBoard.numbers[i][i]
				// nextBoard.wins[11][i] = nextBoard.numbers[i][4-i]
			}
			data.boards = append(data.boards, nextBoard)
			// fmt.Printf("Board output:\n\t%+v\n", nextBoard)
		}
	}
	return data
}

func solution(input fileData) []int {
	// fmt.Printf("%+v\n", input)
	seen := make(map[int]bool)
	won := make(map[int]bool)
	winScores := []int{}

	for _, pick := range input.picks {
		seen[pick] = true
		fmt.Println("Picked:", pick)

		for idx, b := range input.boards {
			if won[idx] {
				continue
			}
			wins := 0
			for _, w := range b.wins {
				totalSeen := 0
				for _, n := range w {
					if seen[n] {
						totalSeen++
					}
				}
				if totalSeen == boardSize {
					wins++
					fmt.Println("\nWINNER", w)
					fmt.Println("board", idx)
					won[idx] = true
				}
			}
			if wins > 0 {
				sumUnseen := 0
				for _, row := range b.numbers {
					for _, n := range row {
						if !seen[n] {
							sumUnseen += n
						}
					}
				}
				// Part1: Uncomment this line
				// return sumUnseen * pick
				fmt.Println("Score:", sumUnseen*pick)
				winScores = append(winScores, sumUnseen*pick)
			}
		}
	}
	return winScores
}
