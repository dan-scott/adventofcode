package day17

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
	"strconv"
	"strings"
)

type day17 struct {
	input string
}

func (d *day17) Open() {
	d.input = inputs.LinesAsString(2021, 17)[0]
}

func (d *day17) Close() {
}

func (d *day17) Part1() string {
	_, _, yS, yE := d.parseInput()
	yIV := 1
	maxY := 0
	chances := 100
Find:
	for {
		yV := yIV
		yC, yM := 0, 0
	Trajectory:
		for {
			if yC+yV < yS {
				yM = 0
				chances--
				if chances < 0 {
					break Find
				} else {
					break Trajectory
				}
			}
			yC += yV
			yV -= 1
			if yC > yM {
				yM = yC
			}
			if yC < yE {
				break Trajectory
			}
		}
		if yM > maxY {
			maxY = yM
		}
		yIV++
	}
	return fmt.Sprint(maxY)
}

func (d *day17) Part2() string {
	xMin, xMax, yMin, yMax := d.parseInput()
	yTimes := make(map[int][]int)
	for yv := yMin; yv < -yMin; yv++ {
		t := 0
		yc := 0
		ycv := yv
		for {
			t++
			yc += ycv
			ycv--
			if yc <= yMax {
				if yc >= yMin {
					if _, ok := yTimes[t]; !ok {
						yTimes[t] = make([]int, 0)
					}
					yTimes[t] = append(yTimes[t], yv)
				} else {
					break
				}
			}
		}
	}
	xTimes := make(map[int][]int)
	xRestTimes := make(map[int][]int)
	for xv := 1; xv <= xMax; xv++ {
		t := 0
		xc := 0
		xcv := xv
		for {
			t++
			xc += xcv
			xcv--
			if xc >= xMin {
				if xc <= xMax {
					if xcv == 0 {
						if _, ok := xRestTimes[t]; !ok {
							xRestTimes[t] = make([]int, 0)
						}
						xRestTimes[t] = append(xRestTimes[t], xv)
						break
					} else {
						if _, ok := xTimes[t]; !ok {
							xTimes[t] = make([]int, 0)
						}
						xTimes[t] = append(xTimes[t], xv)
					}
				} else {
					break
				}
			}
			if xcv <= 0 {
				break
			}
		}
	}
	distinct := make(map[vec2.Vec2]interface{})
	for t, ys := range yTimes {
		if xs, ok := xTimes[t]; ok {
			for _, x := range xs {
				for _, y := range ys {
					distinct[vec2.Of(x, y)] = nil
				}
			}
		}
		for rt, xs := range xRestTimes {
			if t >= rt {
				for _, x := range xs {
					for _, y := range ys {
						distinct[vec2.Of(x, y)] = nil
					}
				}
			}
		}
	}
	return fmt.Sprint(len(distinct))
}

func (d *day17) parseInput() (xS, xE, yS, yE int) {
	parts := strings.Split(d.input[13:], ", ")
	xS, xE = parseMinMax(parts[0])
	yS, yE = parseMinMax(parts[1])
	return
}

func parseMinMax(s string) (min, max int) {
	pts := strings.Split(s[2:], "..")
	min, _ = strconv.Atoi(pts[0])
	max, _ = strconv.Atoi(pts[1])
	return
}

func New() runner.LegacyDay {
	return &day17{}
}
