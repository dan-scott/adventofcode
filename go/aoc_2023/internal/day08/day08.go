package day08

import (
	"bufio"
	"fmt"
	"github.com/danscott/adventofcode/go/common/mathutils"
	"github.com/danscott/adventofcode/go/common/runner"
	"strings"
)

type day08 struct {
}

func (d *day08) Year() uint {
	return 2023
}

func (d *day08) Day() uint {
	return 8
}

func (d *day08) Part1(input *bufio.Scanner) string {
	dirs, nodeMap := parseInput(input)

	steps := 0
	dirsLen := len(dirs)
	n := nodeMap["AAA"]
	for n.id != "ZZZ" {
		switch dirs[steps%dirsLen] {
		case 'L':
			n = n.left
		case 'R':
			n = n.right
		}

		steps += 1
	}
	return fmt.Sprint(steps)
}

func (d *day08) Part2(input *bufio.Scanner) string {
	dirs, nodeMap := parseInput(input)

	var nodes []*node
	for id, n := range nodeMap {
		if id[2] == 'A' {
			nodes = append(nodes, n)
		}
	}

	var counts []int64
	dirsLen := len(dirs)
	var diffs []int64
	for _, n := range nodes {
		steps := 0
		visited := make(map[string]int, 0)
		current := n
		for {
			if _, ok := visited[current.id]; ok {
				steps += 1
				break
			}
			visited[current.id] = steps
			switch dirs[steps%dirsLen] {
			case 'L':
				current = current.left
			case 'R':
				current = current.right
			}
			steps += 1
		}
		diffs = append(diffs, int64(visited[current.id]))
		counts = append(counts, int64(steps-visited[current.id]))
	}

	cycles := mathutils.LCM(counts[0], counts[1], counts[2:]...)

	return fmt.Sprint(cycles)
}

func New() runner.Day {
	return &day08{}
}

func parseInput(input *bufio.Scanner) ([]rune, map[string]*node) {
	input.Scan()
	dirs := []rune(input.Text())
	input.Scan()

	nodeMap := make(map[string]*node, 0)
	getNode := func(id string) *node {
		if n, ok := nodeMap[id]; ok {
			return n
		}
		n := &node{id, nil, nil}
		nodeMap[id] = n
		return n
	}

	for input.Scan() {
		parts := strings.Split(input.Text(), " = ")
		n := getNode(parts[0])
		n.left = getNode(parts[1][1:4])
		n.right = getNode(parts[1][6:9])
	}
	return dirs, nodeMap
}

type node struct {
	id    string
	left  *node
	right *node
}
