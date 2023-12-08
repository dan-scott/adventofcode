package day11

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"github.com/danscott/adventofcode/go/common/vec2"
)

type day11 struct {
	lines []string
}

type cave struct {
	jellies []uint8
	w, h    int
}

func parseCave(lines []string) *cave {
	jellies, w, h := parseLines(lines)
	return &cave{jellies, w, h}
}

func (c *cave) iterate() int {
	next := make([]uint8, len(c.jellies))
	flashed := make([]bool, len(c.jellies))
	toReset := make([]int, 0, len(c.jellies))
	toFlash := make([]int, 0, len(c.jellies))
	for j, e := range c.jellies {
		next[j] = e + 1
		if next[j] > 9 {
			toFlash = append(toFlash, j)
		}
	}
	for len(toFlash) > 0 {
		j := toFlash[0]
		toFlash = toFlash[1:]
		if flashed[j] {
			continue
		}
		flashed[j] = true
		toReset = append(toReset, j)
		for _, nj := range c.neighbours(j) {
			next[nj]++
			if next[nj] > 9 && !flashed[nj] {
				toFlash = append(toFlash, nj)
			}
		}
	}
	for _, j := range toReset {
		next[j] = 0
	}
	c.jellies = next
	return len(toReset)
}

func (d *day11) Open() {
	d.lines = inputs.LinesAsString(2021, 11)
}

func (d *day11) Close() {
	d.lines = nil
}

func (d *day11) Part1() string {
	flashCt := 0
	cave := parseCave(d.lines)
	for iter := 0; iter < 100; iter++ {
		flashCt += cave.iterate()
	}
	return fmt.Sprint(flashCt)
}

func (d *day11) Part2() string {
	cave := parseCave(d.lines)
	iter := 0
	for true {
		if iter > 1000 {
			break
		}
		iter++
		if cave.iterate() == 100 {
			return fmt.Sprint(iter)
		}
	}
	return "Too many iterations..."
}

func parseLines(lines []string) (jellies []uint8, w, h int) {
	w = len(lines[0])
	h = len(lines)
	jellies = make([]uint8, w*h)
	for y, l := range lines {
		for x := 0; x < w; x++ {
			jellies[y*w+x] = l[x] - '0'
		}
	}
	return
}

var neighbourDeltas = []vec2.Vec2{
	vec2.Of(-1, -1),
	vec2.Of(+0, -1),
	vec2.Of(+1, -1),
	vec2.Of(-1, +0),
	vec2.Of(+1, +0),
	vec2.Of(-1, +1),
	vec2.Of(+0, +1),
	vec2.Of(+1, +1),
}

var nCache = make([][]int, 1000)

func (c *cave) neighbours(flash int) []int {
	if nCache[flash] != nil {
		return nCache[flash]
	}
	ns := make([]int, 0, 8)
	fp := vec2.OfIndex(flash, c.w)
	for _, nd := range neighbourDeltas {
		n := fp.Add(nd)
		if c.valid(n) {
			ns = append(ns, n.Index(c.w))
		}
	}
	nCache[flash] = ns
	return ns
}

func (c *cave) valid(v vec2.Vec2) bool {
	return v.X >= 0 && v.X < c.w && v.Y >= 0 && v.Y < c.h
}

func New() runner.LegacyDay {
	return &day11{}
}
