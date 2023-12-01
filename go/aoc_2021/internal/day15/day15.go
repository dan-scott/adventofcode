package day15

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
)

func New() runner.LegacyDay {
	return &day15{}
}

type day15 struct {
	lines []string
}

func (d *day15) Open() {
	d.lines = inputs.LinesAsString(2021, 15)
}

func (d *day15) Close() {
	d.lines = nil
}

func (d *day15) Part1() string {
	return smallestPath(d.parseLines())
}

func smallestPath(lines [][]uint8) string {
	paths := make(map[vec2.Vec2]int)
	paths[vec2.Of(0, 0)] = 0
	end := vec2.Of(len(lines[0])-1, len(lines)-1)
	heads := []vec2.Vec2{vec2.Of(0, 0)}
	checkHead := func(v, x, y int) {
		nh := vec2.Of(x, y)
		nv := v + int(lines[nh.Y][nh.X])
		if ov, ok := paths[nh]; !ok || (ov > nv) {
			paths[nh] = nv
			heads = append(heads, nh)
		}
	}
	for len(heads) > 0 {
		h := heads[0]
		heads = heads[1:]
		cv := paths[h]
		if h.X > 0 {
			checkHead(cv, h.X-1, h.Y)
		}
		if h.X < end.X {
			checkHead(cv, h.X+1, h.Y)
		}
		if h.Y > 0 {
			checkHead(cv, h.X, h.Y-1)
		}
		if h.Y < end.Y {
			checkHead(cv, h.X, h.Y+1)
		}
	}
	return fmt.Sprint(paths[end])
}

func (d *day15) Part2() string {
	lines := d.getBigMap()
	return smallestPath(lines)
}

func (d *day15) parseLines() [][]uint8 {
	nums := make([][]uint8, len(d.lines))
	for y := 0; y < len(d.lines); y++ {
		nums[y] = make([]uint8, len(d.lines[y]))
		for x := 0; x < len(d.lines[y]); x++ {
			nums[y][x] = d.lines[y][x] - '0'
		}
	}
	return nums
}

func (d *day15) getBigMap() [][]uint8 {
	lines := d.parseLines()
	h := len(d.lines)
	w := len(d.lines[0])
	big := make([][]uint8, h*5)
	for i := 0; i < h; i++ {
		big[i] = make([]byte, w*5)
		copy(big[i], lines[i])
	}
	for i := h; i < h*5; i++ {
		big[i] = make([]byte, w*5)
	}
	vs := make([]uint8, 9)
	for i := range vs {
		vs[i] = uint8(i + 1)
	}
	for y := 0; y < h*5; y++ {
		for x := 0; x < w; x++ {
			idx := big[y][x] - 1
			for i := 1; i < 5; i++ {
				idx = (idx + 1) % 9
				big[y][x+i*w] = vs[idx]
				if y < h {
					big[y+i*h][x] = vs[idx]
				}
			}
		}
	}
	return big
}
