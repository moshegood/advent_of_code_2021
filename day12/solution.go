package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	fmt.Println(solution(input, 0))
	fmt.Println(solution(input, 1))
}

type edge struct {
	from, to string
}

type tree struct {
	name     string
	children []*tree
}

func (n *tree) leaves() []*tree {
	if len(n.children) == 0 {
		return []*tree{n}
	}
	list := []*tree{}
	for _, c := range n.children {
		list = append(list, c.leaves()...)
	}
	return list
}

func (root *tree) buildTree(options map[string][]string, smallRepeats int, ancestors []string) {
	ancestors = append(ancestors, root.name)
	for _, option := range options[root.name] {
		// don't revisit lowercase rooms too often
		seen := 0
		if option[0] >= 'a' && option[0] <= 'z' {
			for _, a := range ancestors {
				if a == option {
					seen++
				}
			}
			if seen > smallRepeats {
				continue
			}
		}
		// fmt.Println("Creating", ancestry, option)
		child := tree{option, nil}
		child.buildTree(options, smallRepeats-seen, ancestors)
		root.children = append(root.children, &child)
	}
}

func parseInput(input string) []edge {
	rows := []edge{}
	for _, line := range strings.Fields(input) {
		parts := strings.Split(line, "-")
		rows = append(rows, edge{parts[0], parts[1]})
	}
	return rows
}

func edgesToMap(edges []edge) map[string][]string {
	m := make(map[string][]string)
	for _, e := range edges {
		if e.to != "start" {
			m[e.from] = append(m[e.from], e.to)
		}
		if e.from != "start" {
			m[e.to] = append(m[e.to], e.from)
		}
	}
	delete(m, "end")
	return m
}

func solution(input string, smallRepeats int) int {
	edges := parseInput(input)
	options := edgesToMap(edges)
	fmt.Printf("Options: %+v\n", options)
	root := tree{"start", nil}
	root.buildTree(options, smallRepeats, nil)
	fmt.Printf("Tree: %+v\n", root)

	total := 0
	leaves := root.leaves()
	fmt.Println("Leaves:")
	for _, l := range leaves {
		fmt.Printf("\t%+v\n", *l)
		if l.name == "end" {
			total++
		}
	}

	return total
}
