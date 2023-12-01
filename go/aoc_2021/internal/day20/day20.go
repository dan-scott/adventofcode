package day20

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"gitlab.com/danscott/adventofcode/go/common/vec2"
)

type day20 struct {
	input []string
}

var smp = []vec2.Vec2{
	vec2.Of(-1, -1),
	vec2.Of(0, -1),
	vec2.Of(1, -1),
	vec2.Of(-1, -0),
	vec2.Of(0, -0),
	vec2.Of(1, -0),
	vec2.Of(-1, +1),
	vec2.Of(0, +1),
	vec2.Of(1, +1),
}

func (d *day20) Open() {
	d.input = inputs.LinesAsString(2021, 20)
}

func (d *day20) Close() {
	d.input = nil
}

func (d *day20) Part1() string {
	i := d.parseInput()
	i.iterate()
	i.iterate()
	return fmt.Sprint(len(i.pxMap))
}

func (d *day20) Part2() string {
	im := d.parseInput()
	for i := 0; i < 50; i++ {
		im.iterate()
	}
	return fmt.Sprint(len(im.pxMap))
}

func (i *img) iterate() {
	next := make(map[vec2.Vec2]interface{})

	for x := i.tl.X - 1; x <= i.br.X+1; x++ {
		for y := i.tl.Y - 1; y <= i.br.Y+1; y++ {
			px := vec2.Of(x, y)
			lookup := 0
			for j, d := range smp {
				p := px.Add(d)
				if _, ok := i.pxMap[p]; ok || (i.oobIsLit && i.isOob(p)) {
					lookup |= 1 << (8 - j)
				}
			}
			shouldBeLit := i.bitmap[lookup] == '#'
			if shouldBeLit {
				next[px] = nil
			}
		}
	}
	i.tl.X--
	i.tl.Y--
	i.br.X++
	i.br.Y++
	i.pxMap = next
	if i.oobIsLit {
		for p := range i.pxMap {
			if i.isOob(p) {
				delete(i.pxMap, p)
			}
		}
	}
	i.oobIsLit = i.bitmap[0] == '#' && !i.oobIsLit
}

func (i *img) print() {
	fmt.Println("Out of bounds is lit: ", i.oobIsLit)
	for y := i.tl.Y - 5; y < i.br.Y+5; y++ {
		for x := i.tl.X - 5; x < i.br.X+5; x++ {
			p := vec2.Of(x, y)
			if _, ok := i.pxMap[p]; ok || (i.oobIsLit && i.isOob(p)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (i *img) isOob(p vec2.Vec2) bool {
	return p.X < i.tl.X || p.Y < i.tl.Y || p.X > i.br.X || p.Y > i.br.Y
}

type img struct {
	bitmap   []byte
	pxMap    map[vec2.Vec2]interface{}
	tl, br   vec2.Vec2
	oobIsLit bool
}

func (d *day20) parseInput() *img {
	bitmap := []byte(d.input[0])
	pxMap := make(map[vec2.Vec2]interface{})
	if len(bitmap) != 512 {
		panic("panik")
	}
	tl := vec2.Of(0, 0)
	br := vec2.Of(len(d.input[2])-1, len(d.input)-3)
	for y, line := range d.input[2:] {
		for x, chr := range line {
			if chr == '.' {
				continue
			}
			pxMap[vec2.Of(x, y)] = nil
		}
	}
	return &img{
		bitmap: bitmap,
		pxMap:  pxMap,
		tl:     tl,
		br:     br,
	}
}

func New() runner.LegacyDay {
	return &day20{}
}
