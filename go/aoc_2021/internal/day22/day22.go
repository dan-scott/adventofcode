package day22

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"strconv"
	"strings"
)

type day22 struct {
	input []string
}

func (d *day22) Open() {
	d.input = inputs.LinesAsString(2021, 22)
}

func (d *day22) Close() {
	d.input = nil
}

func (d *day22) Part1() string {
	onMap := make(map[vec3]interface{})
	bounds := cube{
		min: v(-50, -50, -50),
		max: v(50, 50, 50),
	}

	for _, l := range d.input {
		on, c := parseLine(l)
		if !bounds.contains(c) {
			continue
		}
		for x := c.min.x; x <= c.max.x; x++ {
			for y := c.min.y; y <= c.max.y; y++ {
				for z := c.min.z; z <= c.max.z; z++ {
					pt := v(x, y, z)
					if on {
						onMap[pt] = nil
					} else {
						delete(onMap, pt)
					}
				}
			}
		}
	}

	return fmt.Sprint(len(onMap))
}

func (d *day22) Part2() string {
	var cubes []cube
	for _, l := range d.input {
		on, c := parseLine(l)
		var next []cube
		for _, cb := range cubes {
			next = append(next, cb.subtract(c)...)
		}
		if on {
			next = append(next, c)
		}
		cubes = next
		for i, c3 := range cubes[:len(cubes)-1] {
			for j := i + 1; j < len(cubes); j++ {
				c2 := cubes[j]
				if c3.intersects(c2) {
					panic("intersected!")
				}
			}
		}
	}
	onCt := 0
	for _, c := range cubes {
		onCt += c.volume()
	}
	return fmt.Sprint(onCt)
}

type cube struct {
	min, max vec3
}

func (c cube) String() string {
	return fmt.Sprintf("[%s %s]", c.min, c.max)
}

func (c cube) volume() int {
	return c.max.sub(c.min).abs().add(v(1, 1, 1)).prod()
}

func (c cube) intersects(o cube) bool {
	return c.cmpTouches(o, cX) && c.cmpTouches(o, cY) && c.cmpTouches(o, cZ)
}

func (c cube) cmpTouches(o cube, i componentId) bool {
	cMin, cMax := c.cmpRange(i)
	oMin, oMax := o.cmpRange(i)

	if oMax < cMin || oMin > cMax {
		return false
	}

	return true
}

func (c cube) cmpRange(i componentId) (int, int) {
	return c.min.getCmp(i), c.max.getCmp(i)
}

func (c cube) containsPoint(v vec3) (a, x, y, z bool) {
	x = v.x >= c.min.x && v.x <= c.max.x
	y = v.y >= c.min.y && v.y <= c.max.y
	z = v.z >= c.min.z && v.z <= c.max.z
	a = x && y && z
	return
}

func (c cube) contains(o cube) bool {
	ma, _, _, _ := c.containsPoint(o.min)
	mb, _, _, _ := c.containsPoint(o.max)
	return ma && mb
}

func (c cube) subtract(o cube) []cube {
	var cubes []cube
	if o.contains(c) {
		return cubes
	}
	cubes = append(cubes, c)
	if !o.intersects(c) {
		return cubes
	}
	_, minx, miny, minz := c.containsPoint(o.min)
	_, maxx, maxy, maxz := c.containsPoint(o.max)

	if minx {
		cubes = splitCmp(cubes, cX, o.min.x)
	}
	if maxx {
		cubes = splitCmp(cubes, cX, o.max.x+1)
	}
	if miny {
		cubes = splitCmp(cubes, cY, o.min.y)
	}
	if maxy {
		cubes = splitCmp(cubes, cY, o.max.y+1)
	}
	if minz {
		cubes = splitCmp(cubes, cZ, o.min.z)
	}
	if maxz {
		cubes = splitCmp(cubes, cZ, o.max.z+1)
	}

	var filtered []cube
	for _, nc := range cubes {
		if !o.contains(nc) {
			if o.intersects(nc) {
				panic("what")
			}
			filtered = append(filtered, nc)
		}
	}
	return filtered
}

func splitCmp(cubes []cube, cmp componentId, val int) []cube {
	var nextCubes []cube
	lb := val - 1
	for _, c := range cubes {
		wasSplit := false
		if lb >= c.min.getCmp(cmp) && lb <= c.max.getCmp(cmp) {
			nextCubes = append(nextCubes, cube{
				min: c.min,
				max: c.max.setCmp(cmp, lb),
			})
			wasSplit = true
		}
		if val >= c.min.getCmp(cmp) && val <= c.max.getCmp(cmp) {
			nextCubes = append(nextCubes, cube{
				min: c.min.setCmp(cmp, val),
				max: c.max,
			})
			wasSplit = true
		}
		if !wasSplit {
			nextCubes = append(nextCubes, c)
		}
	}
	return nextCubes
}

func parseLine(line string) (on bool, c cube) {
	on = true
	splitAt := 3
	if line[1] == 'f' {
		splitAt = 4
		on = false
	} else {
	}
	pts := strings.Split(line[splitAt:], ",")
	min := vec3{}
	max := vec3{}
	min.x, max.x = getRange(pts[0])
	min.y, max.y = getRange(pts[1])
	min.z, max.z = getRange(pts[2])
	c = cube{min, max}
	return
}

func getRange(inpt string) (min int, max int) {
	pts := strings.Split(inpt[2:], "..")
	min, _ = strconv.Atoi(pts[0])
	max, _ = strconv.Atoi(pts[1])
	if min > max {
		min, max = max, min
	}
	return
}

func New() runner.Day {
	return &day22{}
}
