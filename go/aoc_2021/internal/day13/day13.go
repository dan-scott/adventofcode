package day13

import (
	"bytes"
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type day13 struct {
	lines []string
}

func (d *day13) Open() {
	d.lines = inputs.LinesAsString(2021, 13)
}

func (d *day13) Close() {
	d.lines = nil
}

func (d *day13) Part1() string {
	paper, folds := d.parse()
	for _, f := range folds[:1] {
		switch f.axis {
		case "y":
			for p := range paper {
				if p.Y < f.num {
					continue
				}
				paper[vec2.Of(p.X, 2*f.num-p.Y)] = nil
				delete(paper, p)
			}
		case "x":
			for p := range paper {
				if p.X < f.num {
					continue
				}
				paper[vec2.Of(2*f.num-p.X, p.Y)] = nil
				delete(paper, p)
			}
		}
	}
	return fmt.Sprint(len(paper))
}

func (d *day13) Part2() string {
	paper, folds := d.parse()
	for _, f := range folds {
		switch f.axis {
		case "y":
			for p := range paper {
				if p.Y < f.num {
					continue
				}
				paper[vec2.Of(p.X, 2*f.num-p.Y)] = nil
				delete(paper, p)
			}
		case "x":
			for p := range paper {
				if p.X < f.num {
					continue
				}
				paper[vec2.Of(2*f.num-p.X, p.Y)] = nil
				delete(paper, p)
			}
		}
	}
	return printMap(paper)
}

type fold struct {
	axis string
	num  int
}

func (d *day13) parse() (paper map[vec2.Vec2]interface{}, folds []fold) {
	paper = make(map[vec2.Vec2]interface{}, len(d.lines))
	i := 0
	for d.lines[i] != "" {
		parts := strings.Split(d.lines[i], ",")
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		paper[vec2.OfInt64(x, y)] = nil
		i++
	}
	i++
	folds = make([]fold, 0, len(d.lines)-i)
	for _, l := range d.lines[i:] {
		parts := strings.Split(l[11:], "=")
		num, _ := strconv.ParseInt(parts[1], 10, 64)
		folds = append(folds, fold{
			axis: parts[0],
			num:  int(num),
		})
	}
	return
}

func New() runner.Day {
	return &day13{}
}

func printMap(m map[vec2.Vec2]interface{}) string {
	max := vec2.Of(0, 0)
	for p := range m {
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	var b bytes.Buffer
	b.WriteByte('\n')
	for y := 0; y <= max.Y; y++ {
		for x := 0; x <= max.X; x++ {
			if _, ok := m[vec2.Of(x, y)]; ok {
				b.WriteByte('@')
			} else {
				b.WriteByte('_')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}
