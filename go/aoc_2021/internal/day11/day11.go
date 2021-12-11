package day11

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
)

type day11 struct {
	lines []string
	w, h  int64
}

func (d *day11) Open() {
	d.lines = inputs.LinesAsString(2021, 11)
	d.w = int64(len(d.lines[0]))
	d.h = int64(len(d.lines))
}

func (d *day11) Close() {
	d.lines = nil
}

func (d *day11) Part1() string {
	current := d.parseLines()
	flashCt := 0
	for iter := 0; iter < 100; iter++ {
		flashed := make(map[vec2.Vec2]bool, len(current))
		flashes := make([]vec2.Vec2, 0, len(current))
		next := make(map[vec2.Vec2]uint8, len(current))
		for p, v := range current {
			next[p] = v + 1
			if v == 9 {
				flashes = append(flashes, p)
			}
		}
		for len(flashes) > 0 {
			f := flashes[0]
			flashes = flashes[1:]
			if flashed[f] {
				continue
			}
			flashed[f] = true
			flashCt++
			for _, p := range d.neighbours(f) {
				next[p]++
				if next[p] > 9 && !flashed[p] {
					flashes = append(flashes, p)
				}
			}
		}
		for p := range flashed {
			next[p] = 0
		}
		current = next
	}
	return fmt.Sprint(flashCt)
}

func (d *day11) Part2() string {
	current := d.parseLines()
	iter := 0
	for true {
		iter++
		flashCt := 0
		flashed := make(map[vec2.Vec2]bool, len(current))
		flashes := make([]vec2.Vec2, 0, len(current))
		next := make(map[vec2.Vec2]uint8, len(current))
		for p, v := range current {
			next[p] = v + 1
			if v == 9 {
				flashes = append(flashes, p)
			}
		}
		for len(flashes) > 0 {
			f := flashes[0]
			flashes = flashes[1:]
			if flashed[f] {
				continue
			}
			flashed[f] = true
			flashCt++
			for _, p := range d.neighbours(f) {
				next[p]++
				if next[p] > 9 && !flashed[p] {
					flashes = append(flashes, p)
				}
			}
		}
		for p := range flashed {
			next[p] = 0
		}
		current = next
		if flashCt == 100 {
			return fmt.Sprint(iter)
		}
	}
	return ""
}

func (d *day11) parseLines() map[vec2.Vec2]uint8 {
	startGrid := make(map[vec2.Vec2]uint8, len(d.lines)*len(d.lines[0]))
	for y, l := range d.lines {
		for x := 0; x < len(d.lines[0]); x++ {
			startGrid[vec2.OfInt(x, y)] = l[x] - '0'
		}
	}
	return startGrid
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

func (d *day11) neighbours(flash vec2.Vec2) []vec2.Vec2 {
	ns := make([]vec2.Vec2, 0, 8)
	for _, nd := range neighbourDeltas {
		n := flash.Add(nd)
		if d.valid(n) {
			ns = append(ns, n)
		}
	}
	return ns
}

func (d *day11) valid(v vec2.Vec2) bool {
	return v.X >= 0 && v.X < d.w && v.Y >= 0 && v.Y < d.h
}

func New() runner.Day {
	return &day11{}
}
