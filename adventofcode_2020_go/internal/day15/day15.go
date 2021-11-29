package day15

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/inputs"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"strconv"
	"strings"
)

type day15 struct {
	init []int
}

func (d *day15) Open() {
	d.load(inputs.LinesAsString(15)[0])
}

func (d *day15) load(line string) {
	vs := strings.Split(line, ",")
	init := make([]int, len(vs))
	for i, n := range vs {
		v, _ := strconv.ParseInt(n, 10, 64)
		init[i] = int(v)
	}
	d.init = init
}

func (d *day15) Close() {
	d.init = nil
}

func (d *day15) Part1() string {
	return d.callForTurn(2020)
}

func (d *day15) Part2() string {
	return d.callForTurn(30000000)
}

func (d *day15) callForTurn(maxTurns int) string {
	called := make([]bool, maxTurns)
	lastCall := make([]int, maxTurns)
	last := 0
	next := 0

	for turn := 0; turn < maxTurns; turn++ {
		if turn < len(d.init) {
			next = d.init[turn]
		} else {
			if !called[last] {
				next = 0
			} else {
				next = turn - lastCall[last]
			}
		}
		if turn > 0 {
			called[last] = true
			lastCall[last] = turn
		}
		last = next
	}

	return fmt.Sprint(last)
}

func New() internal.Day {
	return &day15{}
}
