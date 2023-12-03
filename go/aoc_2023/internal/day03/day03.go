package day03

import (
	"bufio"
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
	"strconv"
)

type day03 struct {
}

func (d *day03) Year() uint {
	return 2023
}

func (d *day03) Day() uint {
	return 3
}

func (d *day03) Part1(input *bufio.Scanner) string {
	nums, syms, _ := parseToMap(input)
	total := 0
	for _, np := range nums {
		for nb := range np.nbs {
			if _, ok := syms[nb]; ok {
				total += np.n
			}
		}
	}
	return fmt.Sprint(total)
}

func (d *day03) Part2(input *bufio.Scanner) string {
	_, syms, nums := parseToMap(input)
	total := 0
	for p, s := range syms {
		if s != '*' || len(nums[p]) != 2 {
			continue
		}
		total += nums[p][0] * nums[p][1]
	}
	return fmt.Sprint(total)
}

func New() runner.Day {
	return &day03{}
}

type numPos struct {
	n   int
	pos []vec2.Vec2
	nbs map[vec2.Vec2]interface{}
}

func parseToMap(input *bufio.Scanner) (numbers []numPos, symbols map[vec2.Vec2]rune, numMap map[vec2.Vec2][]int) {
	numbers = make([]numPos, 0)
	symbols = make(map[vec2.Vec2]rune)
	numMap = make(map[vec2.Vec2][]int)
	y := -1
	for input.Scan() {
		y += 1
		line := input.Text()
		inNum := false
		numStart := 0
		for x, r := range []rune(line) {
			d := isDigit(r)
			if inNum && !d {
				n, _ := strconv.ParseInt(line[numStart:x], 10, 32)
				np := numPos{
					n:   int(n),
					pos: []vec2.Vec2{},
					nbs: make(map[vec2.Vec2]interface{}),
				}
				for xi := numStart; xi < x; xi++ {
					nextPos := vec2.Of(xi, y)
					np.pos = append(np.pos, nextPos)
					for _, p := range nextPos.Neighbours() {
						np.nbs[p] = nil
					}
				}
				for p, _ := range np.nbs {
					if _, ok := numMap[p]; ok {
						numMap[p] = append(numMap[p], np.n)
					} else {
						numMap[p] = []int{np.n}
					}
				}
				numbers = append(numbers, np)
				inNum = false
			}

			if !inNum && d {
				inNum = true
				numStart = x
			}

			if r != '.' && !d {
				symbols[vec2.Of(x, y)] = r
			}
		}
		if inNum {
			n, _ := strconv.ParseInt(line[numStart:], 10, 32)
			np := numPos{
				n:   int(n),
				pos: []vec2.Vec2{},
				nbs: make(map[vec2.Vec2]interface{}),
			}
			for xi := numStart; xi < len(line); xi++ {
				nextPos := vec2.Of(xi, y)
				np.pos = append(np.pos, nextPos)
				for _, p := range nextPos.Neighbours() {
					np.nbs[p] = nil
				}
			}
			numbers = append(numbers, np)
			inNum = false
		}

	}
	return
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
