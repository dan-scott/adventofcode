package day13

import (
	"bytes"
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"github.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type day13 struct {
	lines []string
}

type paper = map[vec2.Vec2]interface{}

type fold struct {
	axis string
	num  int
}

func (d *day13) Open() {
	d.lines = inputs.LinesAsString(2021, 13)
}

func (d *day13) Close() {
	d.lines = nil
}

func (d *day13) Part1() string {
	p, folds := d.parse()
	p = doFold(p, folds[0])
	return fmt.Sprint(len(p))
}

func (d *day13) Part2() string {
	p, folds := d.parse()
	for _, f := range folds {
		p = doFold(p, f)
	}
	return printMap(p)
}

func (d *day13) parse() (p paper, folds []fold) {
	p = make(paper, len(d.lines))
	i := 0
	for d.lines[i] != "" {
		parts := strings.Split(d.lines[i], ",")
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		p[vec2.OfInt64(x, y)] = nil
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
func doFold(p paper, f fold) paper {
	switch f.axis {
	case "y":
		for dot := range p {
			if dot.Y < f.num {
				continue
			}
			p[vec2.Of(dot.X, 2*f.num-dot.Y)] = nil
			delete(p, dot)
		}
	case "x":
		for dot := range p {
			if dot.X < f.num {
				continue
			}
			p[vec2.Of(2*f.num-dot.X, dot.Y)] = nil
			delete(p, dot)
		}
	}
	return p
}

func New() runner.LegacyDay {
	return &day13{}
}

func printMap(m paper) string {
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
