package day05

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"github.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type grad = uint8

const (
	HORIZONTAL grad = iota
	NEGATIVE
	VERTICAL
	POSITIVE
)

type day05 struct {
	input []string
}

type line struct {
	minX, minY, maxX, maxY int
	g                      grad
}

func (d *day05) Open() {
	d.loadLines(inputs.LinesAsString(2021, 5))
}

func (d *day05) Close() {
	d.input = nil
}

func (d *day05) Part1() string {
	lines := d.getLines()
	plot := make([]int, 1000000)
	ct := 0
	inc := func(id int) {
		plot[id]++
		if plot[id] == 2 {
			ct++
		}
	}
	for _, l := range lines {
		if l.g == HORIZONTAL {
			min := l.minX + l.minY*1000
			max := l.maxY*1000 + l.maxX
			for idx := min; idx <= max; idx++ {
				inc(idx)
			}
		} else if l.g == VERTICAL {
			max := l.maxY*1000 + l.minX
			for idx := l.minY*1000 + l.minX; idx <= max; idx += 1000 {
				inc(idx)
			}
		}
	}
	return fmt.Sprint(ct)
}

func grd(a, b vec2.Vec2) int {
	return (a.X - b.X) / (a.Y - b.Y)
}

func (d *day05) Part2() string {
	lines := d.getLines()
	plot := make([]int, 1000000)
	ct := 0
	inc := func(id int) {
		plot[id]++
		if plot[id] == 2 {
			ct++
		}
	}
	for _, l := range lines {
		switch l.g {
		case HORIZONTAL:
			y := l.minY * 1000
			for x := l.minX; x <= l.maxX; x++ {
				inc(x + y)
			}
		case VERTICAL:
			for y := l.minY * 1000; y <= l.maxY*1000; y += 1000 {
				inc(l.minX + y)
			}
		case POSITIVE:
			x := l.minX
			for y := l.minY * 1000; y <= l.maxY*1000; y += 1000 {
				inc(x + y)
				x++
			}
		case NEGATIVE:
			x := l.maxX
			for y := l.minY * 1000; y <= l.maxY*1000; y += 1000 {
				inc(x + y)
				x--
			}
		}
	}
	return fmt.Sprint(ct)
}

func (d *day05) loadLines(input []string) {
	d.input = input
}

func parseVec(s string) vec2.Vec2 {
	ns := strings.Split(s, ",")
	x, _ := strconv.ParseInt(ns[0], 10, 64)
	y, _ := strconv.ParseInt(ns[1], 10, 64)
	return vec2.OfInt64(x, y)
}

func (d *day05) getLines() []line {
	lines := make([]line, len(d.input))
	for i, l := range d.input {
		a, b := parseLine(l)
		minX, maxX, minY, maxY := a.X, b.X, a.Y, b.Y
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		if minY > maxY {
			minY, maxY = maxY, minY
		}
		ln := line{
			minX: minX,
			minY: minY,
			maxX: maxX,
			maxY: maxY,
			g:    HORIZONTAL,
		}
		if a.Y != b.Y {
			ln.g = grad(grd(a, b) + 2)
		}
		lines[i] = ln
	}
	return lines
}

func parseLine(l string) (vec2.Vec2, vec2.Vec2) {
	p := strings.Split(l, " -> ")
	return parseVec(p[0]), parseVec(p[1])
}

func New() runner.LegacyDay {
	return &day05{}
}
