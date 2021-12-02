package day10

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"sort"
)

type day10 struct {
	adapters []int
}

func (d *day10) Open() {
	d.loadLines(inputs.LinesAsInt(2020, 10))
}
func (d *day10) loadLines(lines []int) {
	d.adapters = lines
	sort.Slice(d.adapters, func(i, j int) bool {
		return d.adapters[i] < d.adapters[j]
	})
}

func (d *day10) Close() {
	d.adapters = nil
}

func (d *day10) Part1() string {
	jolts := 0
	dist := make([]int, 4)
	dist[3]++
	for _, a := range d.adapters {
		dist[a-jolts]++
		jolts = a
	}
	return fmt.Sprint(dist[1] * dist[3])
}

func (d *day10) Part2() string {
	pathCt := make(map[int]int, len(d.adapters))
	for i, a := range d.adapters {
		if a <= 3 {
			pathCt[a]++
		}
		for _, pa := range d.adapters[:i] {
			if a-pa <= 3 {
				pathCt[a] += pathCt[pa]
			}
		}
	}
	return fmt.Sprint(pathCt[d.adapters[len(d.adapters)-1]])
}

func New() runner.Day {
	return &day10{}
}
