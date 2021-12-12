package day12

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"strings"
)

type day12 struct {
	input []string
}

type cave struct {
	big   bool
	id    string
	conns map[string]*cave
}

func (d *day12) Open() {
	d.input = inputs.LinesAsString(2021, 12)
}

func (d *day12) Close() {
	d.input = nil
}

func (d *day12) Part1() string {
	fullPaths := make([][]*cave, 0)
	queue := [][]*cave{{d.getStart()}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastCave := path[len(path)-1]
	Gen:
		for _, c := range lastCave.conns {
			if c.id == "end" {
				fullPaths = append(fullPaths, append(path, c))
				continue Gen
			}
			if !c.big {
				for _, p := range path {
					if p.id == c.id {
						continue Gen
					}
				}
			}
			nextPath := make([]*cave, len(path), len(path)+1)
			copy(nextPath, path)
			nextPath = append(nextPath, c)
			queue = append(queue, nextPath)
		}
	}

	return fmt.Sprint(len(fullPaths))
}

type cavePath struct {
	doubled bool
	caves   []*cave
}

func (p *cavePath) connections() map[string]*cave {
	return p.caves[len(p.caves)-1].conns
}

func (p *cavePath) tryAdd(c *cave) (*cavePath, bool) {
	hasCave := false
	if !c.big {
		hasCave = p.hasCave(c)
	}
	if p.doubled && !c.big && hasCave {
		return nil, false
	}
	newCaves := make([]*cave, len(p.caves), len(p.caves)+1)
	copy(newCaves, p.caves)
	newCaves = append(newCaves, c)
	return &cavePath{
		doubled: p.doubled || hasCave,
		caves:   newCaves,
	}, true
}

func (p *cavePath) hasCave(c *cave) bool {
	for _, pc := range p.caves {
		if pc.id == c.id {
			return true
		}
	}

	return false
}

func (d *day12) Part2() string {
	fullPaths := make([]*cavePath, 0)
	queue := []*cavePath{{
		doubled: false,
		caves:   []*cave{d.getStart()},
	}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, c := range p.connections() {
			if c.id == "start" {
				continue
			}
			next, ok := p.tryAdd(c)
			if !ok {
				continue
			}
			if c.id == "end" {
				fullPaths = append(fullPaths, next)
				continue
			}
			queue = append(queue, next)
		}
	}
	return fmt.Sprint(len(fullPaths))
}

func (d *day12) getStart() *cave {
	caves := make(map[string]*cave, len(d.input)*2)
	getCave := func(id string) *cave {
		c, cached := caves[id]
		if !cached {
			c = &cave{
				big:   int(id[0])-int('a') < 0,
				id:    id,
				conns: make(map[string]*cave, 0),
			}
			caves[id] = c
		}
		return c
	}
	for _, l := range d.input {
		parts := strings.Split(l, "-")
		left, right := getCave(parts[0]), getCave(parts[1])
		if _, ok := left.conns[right.id]; !ok {
			left.conns[right.id] = right
		}
		if _, ok := right.conns[left.id]; !ok {
			right.conns[left.id] = left
		}
	}
	return caves["start"]
}

func New() runner.Day {
	return &day12{}
}
