package day05

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type day05 struct {
	lines    [][]vec2.Vec2
	min, max vec2.Vec2
}

func (d *day05) Open() {
	d.loadLines(inputs.LinesAsString(2021, 5))
}

func (d *day05) Close() {
	d.lines = nil
}

func (d *day05) Part1() string {
	plot := make([][]int, 1000)
	for i := d.min.X; i <= d.max.X; i++ {
		plot[i] = make([]int, 1000)
	}
	ct := 0
	for _, l := range d.lines {
		if l[0].X == l[1].X {
			for y := minY(l); y <= maxY(l); y++ {
				plot[l[0].X][y]++

				if plot[l[0].X][y] == 2 {
					ct++
				}
			}
		} else if l[0].Y == l[1].Y {
			for x := minX(l); x <= maxX(l); x++ {
				plot[x][l[0].Y]++
				if plot[x][l[0].Y] == 2 {
					ct++
				}
			}
		}
	}
	return fmt.Sprint(ct)
}

func maxY(line []vec2.Vec2) int {
	if line[0].Y > line[1].Y {
		return line[0].Y
	}
	return line[1].Y
}

func minY(line []vec2.Vec2) int {
	if line[0].Y < line[1].Y {
		return line[0].Y
	}
	return line[1].Y
}

func maxX(line []vec2.Vec2) int {
	if line[0].X > line[1].X {
		return line[0].X
	}
	return line[1].X
}

func minX(line []vec2.Vec2) int {
	if line[0].X < line[1].X {
		return line[0].X
	}
	return line[1].X
}

func gradient(line []vec2.Vec2) int {
	return (line[1].X - line[0].X) / (line[1].Y - line[0].Y)
}

func (d *day05) Part2() string {
	plot := make([][]int, 1000)
	for i := d.min.X; i <= d.max.X; i++ {
		plot[i] = make([]int, 1000)
	}
	ct := 0
	for _, l := range d.lines {
		if l[0].X == l[1].X {
			for y := minY(l); y <= maxY(l); y++ {
				plot[l[0].X][y]++
				if plot[l[0].X][y] == 2 {
					ct++
				}
			}
		} else if l[0].Y == l[1].Y {
			for x := minX(l); x <= maxX(l); x++ {
				plot[x][l[0].Y]++
				if plot[x][l[0].Y] == 2 {
					ct++
				}
			}
		} else if gradient(l) > 0 {
			x := minX(l)
			for y := minY(l); y <= maxY(l); y++ {
				plot[x][y]++
				if plot[x][y] == 2 {
					ct++
				}
				x++
			}
		} else {
			x := minX(l)
			for y := maxY(l); y >= minY(l); y-- {
				plot[x][y]++
				if plot[x][y] == 2 {
					ct++
				}
				x++
			}
		}
	}
	return fmt.Sprint(ct)
}

func (d *day05) loadLines(input []string) {
	lines := make([][]vec2.Vec2, len(input))
	min := vec2.Of(1000, 1000)
	max := vec2.Of(-1, -1)
	for i, l := range input {
		pts := strings.Split(l, " -> ")
		lines[i] = []vec2.Vec2{parseVec(pts[0]), parseVec(pts[1])}
		if lines[i][0].X < min.X {
			min.X = lines[i][0].X
		}
		if lines[i][1].X < min.X {
			min.X = lines[i][1].X
		}
		if lines[i][0].X > max.X {
			max.X = lines[i][0].X
		}
		if lines[i][1].X > max.X {
			max.X = lines[i][1].X
		}
		if lines[i][0].Y < min.Y {
			min.Y = lines[i][0].Y
		}
		if lines[i][1].Y < min.Y {
			min.Y = lines[i][1].Y
		}
		if lines[i][0].Y > max.Y {
			max.Y = lines[i][0].Y
		}
		if lines[i][1].Y > max.Y {
			max.Y = lines[i][1].Y
		}
	}
	d.min = min
	d.max = max
	d.lines = lines
}

func parseVec(s string) vec2.Vec2 {
	ns := strings.Split(s, ",")
	x, _ := strconv.ParseInt(ns[0], 10, 64)
	y, _ := strconv.ParseInt(ns[1], 10, 64)
	return vec2.OfInt64(x, y)
}

func New() runner.Day {
	return &day05{}
}
