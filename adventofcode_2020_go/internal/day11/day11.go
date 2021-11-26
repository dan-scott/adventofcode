package day11

import (
	"bytes"
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/inputs"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"strings"
)

type day11 struct {
	deck *deck
}

type deck struct {
	cells     []byte
	w, h, pop int
	ray       bool
}

func newDeck(lines []string) *deck {
	return &deck{
		cells: []byte(strings.Join(lines, "")),
		w:     len(lines[0]),
		h:     len(lines),
		pop:   0,
	}
}

func (d *deck) next() *deck {
	next := &deck{
		cells: make([]byte, len(d.cells)),
		w:     d.w,
		h:     d.h,
		ray:   d.ray,
	}
	maxN := 4
	if d.ray {
		maxN = 5
	}
	for i, c := range d.cells {
		if c == '.' {
			next.setCell(i, '.')
			continue
		}
		nc := d.countOccupiedNeighbours(i)
		if c == 'L' && nc == 0 {
			next.setCell(i, '#')
		} else if c == '#' && nc >= maxN {
			next.setCell(i, 'L')
		} else {
			next.setCell(i, c)
		}
	}
	return next
}

func (d *deck) setCell(i int, c byte) {
	d.cells[i] = c
	if c == '#' {
		d.pop++
	}
}

func (d *deck) equals(o *deck) bool {
	return bytes.Compare(d.cells, o.cells) == 0
}

func (d *deck) String() string {
	s := ""
	for i := 0; i < d.h; i++ {
		s += string(d.cells[i*d.w : (i+1)*d.w])
		if i < d.h-1 {
			s += "\n"
		}
	}

	return s
}

type pt struct {
	x, y int
}

func (p pt) add(o pt) pt {
	return pt{
		p.x + o.x,
		p.y + o.y,
	}
}

var deltas = []pt{
	{-1, -1},
	{+0, -1},
	{+1, -1},
	{-1, +0},
	{+1, +0},
	{-1, +1},
	{+0, +1},
	{+1, +1},
}

func (p pt) neighbours() []pt {
	return []pt{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

func (d *deck) ptAt(i int) pt {
	return pt{
		i % d.w,
		i / d.w,
	}
}

func (d *deck) isPtValid(p pt) bool {
	return p.x >= 0 && p.x < d.w && p.y >= 0 && p.y < d.h
}

func (d *deck) cellAtPt(p pt) byte {
	return d.cells[p.x+p.y*d.w]
}

func (d *deck) countOccupiedNeighbours(i int) int {
	if d.ray {
		return d.countRayNeighbours(i)
	}
	total := 0
	for _, n := range d.ptAt(i).neighbours() {
		if d.isPtValid(n) && d.cellAtPt(n) == '#' {
			total++
		}
	}
	return total
}

func (d *deck) countRayNeighbours(i int) int {
	total := 0
	pt := d.ptAt(i)
	for _, delta := range deltas {
		n := pt.add(delta)
	CastRay:
		for d.isPtValid(n) {
			switch d.cellAtPt(n) {
			case '#':
				total++
				break CastRay
			case 'L':
				break CastRay
			}
			if d.cellAtPt(n) == '#' {
				total++
				break CastRay
			}
			n = n.add(delta)
		}
	}
	return total
}

func (d *day11) Open() {
	d.loadLines(inputs.LinesAsString(11))
}

func (d *day11) loadString(s string) {
	d.loadLines(strings.Split(s, "\n"))
}

func (d *day11) loadLines(lines []string) {
	d.deck = newDeck(lines)
}

func (d *day11) Close() {
	d.deck = nil
}

func (d *day11) Part1() string {
	stable := false
	current := d.deck
	for !stable {
		next := current.next()
		stable = current.equals(next)
		current = next
	}
	return fmt.Sprint(current.pop)
}

func (d *day11) Part2() string {
	stable := false
	d.deck.ray = true
	current := d.deck
	for !stable {
		next := current.next()
		stable = current.equals(next)
		current = next
	}
	return fmt.Sprint(current.pop)
}

func New() internal.Day {
	return &day11{}
}
