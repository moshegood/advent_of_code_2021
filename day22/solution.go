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
	fmt.Println(part1(input, 50))
	fmt.Println(part2(input))
}

type rng struct {
	min, max int
}

func (s rng) Size() int {
	if s.max < s.min {
		return 0
	}
	return s.max - s.min + 1
}

func (s rng) Intersect(other rng) rng {
	if s.max > other.max {
		s.max = other.max
	}
	if s.min < other.min {
		s.min = other.min
	}
	return s
}

var emptyRng rng = rng{0, -1}

type cube struct {
	xs rng
	ys rng
	zs rng
}

var emptyCube = cube{emptyRng, emptyRng, emptyRng}

func (c cube) Size() int {
	return c.xs.Size() * c.ys.Size() * c.zs.Size()
}

func (i cube) Intersect(j cube) cube {
	i.xs = i.xs.Intersect(j.xs)
	i.ys = i.ys.Intersect(j.ys)
	i.zs = i.zs.Intersect(j.zs)
	return i
}

type step struct {
	on bool
	cube
}

func parseInput(input string) []step {
	rows := []step{}
	for _, line := range strings.Split(input, "\n") {
		instr := step{}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		if parts[0] == "on" {
			instr.on = true
		}
		dimensions := strings.Split(parts[1], ",")
		for i, dim := range dimensions {
			minMax := strings.Split(dim[2:], "..")
			dimRng := rng{}
			dimRng.min, _ = strconv.Atoi(minMax[0])
			dimRng.max, _ = strconv.Atoi(minMax[1])
			switch i {
			case 0:
				instr.xs = dimRng
			case 1:
				instr.ys = dimRng
			case 2:
				instr.zs = dimRng
			default:
				panic("too many dimensions!!")
			}
		}
		rows = append(rows, instr)
	}
	return rows
}

func part1(input string, max int) int {
	steps := parseInput(input)
	fmt.Printf("Instructions: %+v\n", steps)

	on := map[[3]int]bool{}
	for _, step := range steps {
		if step.xs.max > max || step.ys.max > max || step.zs.max > max {
			continue
		}
		if step.xs.min < -max || step.ys.min < -max || step.zs.min < -max {
			continue
		}
		for x := step.xs.min; x <= step.xs.max; x++ {
			for y := step.ys.min; y <= step.ys.max; y++ {
				for z := step.zs.min; z <= step.zs.max; z++ {
					on[[3]int{x, y, z}] = step.on
				}
			}
		}
	}

	count := 0
	for _, v := range on {
		if v {
			count++
		}
	}
	return count
}

// Recall that union(A,B) is written A ∪ B.
// Recall that intersections(A,B) is written: A ∩ B.
// Recall that |A ∪ B| = |A| + |B| - |A ∩ B|.
// Expanding on that, if B = C ∪ D then:
// |A ∪ B|
// = |A ∪ C ∪ D|
// = |A| + |C| + |D| - |A ∩ C| - |A ∩ D| - |C ∩ D| + |A ∩ C ∩ D|.
func part2(input string) int {
	steps := parseInput(input)
	fmt.Printf("Instructions: %+v\n", steps)

	count := 0
	stepsTaken := []step{}
	for _, this := range steps {
		newSteps := []step{}
		if this.on {
			count += this.Size()
			newSteps = append(newSteps, this)
		}
		for _, that := range stepsTaken {
			overlap := this.cube.Intersect(that.cube)
			if overlap.Size() == 0 {
				continue
			}
			newStep := step{!that.on, overlap}
			newSteps = append(newSteps, newStep)
			if newStep.on {
				count += newStep.Size()
			} else {
				count -= newStep.Size()
			}
		}
		stepsTaken = append(stepsTaken, newSteps...)
		fmt.Printf("Found %+v. Count now %d.\n", this, count)
	}

	return count
}
