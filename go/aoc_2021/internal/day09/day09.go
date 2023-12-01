package day09

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
	"sort"
)

type day09 struct {
	grid          [][]uint8
	height, width int
}

func (d *day09) Open() {
	s, closeScanner := inputs.Scanner(2021, 9)
	defer closeScanner()
	d.grid = make([][]uint8, 0, 100)
	for s.Scan() {
		l := s.Text()
		row := make([]uint8, len(l))
		for i := 0; i < len(l); i++ {
			row[i] = l[i] - '0'
		}
		d.grid = append(d.grid, row)
	}

	d.height = len(d.grid) - 1
	d.width = len(d.grid[0]) - 1
}

func (d *day09) Close() {
	d.grid = nil
}

func (d *day09) Part1() string {
	sum := 0
	for y, row := range d.grid {
	CellCheck:
		for x, val := range row {
			if d.notLocalMin(x, y) {
				continue CellCheck
			}
			sum += int(val) + 1
		}
	}
	return fmt.Sprint(sum)
}

func (d *day09) Part2() string {
	basins := make([]int, 0)
	for y, row := range d.grid {
	CellCheck:
		for x := range row {
			if d.notLocalMin(x, y) {
				continue CellCheck
			}
			basins = append(basins, d.calcFloodFill(x, y))
		}
	}

	sort.SliceStable(basins, func(i int, j int) bool {
		return basins[i] > basins[j]
	})

	return fmt.Sprint(basins[0] * basins[1] * basins[2])
}

func (d *day09) notLocalMin(x, y int) bool {
	val := d.grid[y][x]
	return (x > 0 && d.grid[y][x-1] <= val) || (x < d.width && d.grid[y][x+1] <= val) || (y > 0 && d.grid[y-1][x] <= val) || (y < d.height && d.grid[y+1][x] <= val)
}

func (d *day09) calcFloodFill(x int, y int) int {
	queue := make([]vec2.Vec2, 1)
	queue[0] = vec2.Of(x, y)
	visited := make([]bool, (d.width+1)*(d.height+1))
	size := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		idx := n.Index(d.width)
		if visited[idx] {
			continue
		}
		visited[idx] = true
		if d.grid[n.Y][n.X] == 9 {
			continue
		}
		size++
		if n.X > 0 {
			queue = append(queue, vec2.Of(n.X-1, n.Y))
		}
		if n.X < d.width {
			queue = append(queue, vec2.Of(n.X+1, n.Y))
		}
		if n.Y > 0 {
			queue = append(queue, vec2.Of(n.X, n.Y-1))
		}
		if n.Y < d.height {
			queue = append(queue, vec2.Of(n.X, n.Y+1))
		}
	}
	return size
}

func New() runner.LegacyDay {
	return &day09{}
}
