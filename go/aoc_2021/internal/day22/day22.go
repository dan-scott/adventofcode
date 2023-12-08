package day22

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
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
	var cubes []toggleCube
	for _, l := range d.input {
		on, c := parseLine(l)

		ct := len(cubes)
		for i := 0; i < ct; i++ {
			cb := cubes[i]
			if o, overlaps := cb.overlappingCube(c); overlaps {
				cubes = append(cubes, toggleCube{o, !cb.on})
			}
		}

		if on {
			cubes = append(cubes, toggleCube{c, true})
		}
	}

	onCt := 0
	for _, c := range cubes {
		if c.on {
			onCt += c.volume()
		} else {
			onCt -= c.volume()
		}
	}
	return fmt.Sprint(onCt)
}

type toggleCube struct {
	cube
	on bool
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

func New() runner.LegacyDay {
	return &day22{}
}
