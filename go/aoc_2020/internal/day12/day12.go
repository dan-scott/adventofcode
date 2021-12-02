package day12

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"math"
	"strconv"
)

const (
	L = 'L'
	R = 'R'
	F = 'F'
)

var dirs = []vec{
	{+1, +0}, // E
	{+0, +1}, // S
	{-1, +0}, // W
	{+0, -1}, // N
}

type vec struct {
	x, y int
}

func (v vec) add(o vec) vec {
	return vec{
		v.x + o.x,
		v.y + o.y,
	}
}

func (v vec) mul(mag int) vec {
	return vec{
		v.x * mag,
		v.y * mag,
	}
}

func (v vec) String() string {
	return fmt.Sprintf("{%d %d}", v.x, v.y)
}

var dirMap map[byte]vec

func init() {
	dirMap = make(map[byte]vec, 4)
	dirMap['E'] = dirs[0]
	dirMap['S'] = dirs[1]
	dirMap['W'] = dirs[2]
	dirMap['N'] = dirs[3]
}

type day12 struct {
	ops []string
	pos vec
	wp  vec
	dir int
}

func (d *day12) Open() {
	d.loadLines(inputs.LinesAsString(2020, 12))
}

func (d *day12) loadLines(lines []string) {
	d.ops = lines
}

func (d *day12) resetPos() {
	d.pos = vec{}
	d.wp = vec{10, -1}
	d.dir = 0
}

func (d *day12) Close() {
	d.ops = nil
}

func (d *day12) Part1() string {
	d.resetPos()
	for _, op := range d.ops {
		val, _ := strconv.ParseInt(op[1:], 10, 64)
		if dir, ok := dirMap[op[0]]; ok {
			d.pos = d.pos.add(dir.mul(int(val)))
		} else {
			switch op[0] {
			case L:
				d.dir = (d.dir + (4 - int(val/90))) % 4
				break
			case R:
				d.dir = (d.dir + int(val/90)) % 4
				break
			case F:
				d.pos = d.pos.add(dirs[d.dir].mul(int(val)))
				break
			}
		}
	}

	return fmt.Sprint(math.Abs(float64(d.pos.x)) + math.Abs(float64(d.pos.y)))
}

func (d *day12) Part2() string {
	d.resetPos()
	for _, op := range d.ops {
		val, _ := strconv.ParseInt(op[1:], 10, 64)
		if dir, ok := dirMap[op[0]]; ok {
			d.wp = d.wp.add(dir.mul(int(val)))
		} else {
			switch op[0] {
			case L:
				d.rotateWp(-val)
				break
			case R:
				d.rotateWp(val)
				break
			case F:
				d.pos = d.pos.add(d.wp.mul(int(val)))
				break
			}
		}
	}
	return fmt.Sprint(math.Abs(float64(d.pos.x)) + math.Abs(float64(d.pos.y)))
}

func (d *day12) rotateWp(deg int64) {
	for i := 0; i < int(((deg+360)/90)%4); i++ {
		d.wp = vec{-d.wp.y, d.wp.x}
	}
}

func New() runner.Day {
	return &day12{}
}
