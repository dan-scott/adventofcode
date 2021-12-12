package day12

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"sort"
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

func (c *cave) String() string {
	return c.id
}

type node struct {
	path    int
	current int
	doubled bool
}

func (d *day12) Open() {
	d.input = inputs.LinesAsString(2021, 12)
}

func (d *day12) Close() {
	d.input = nil
}

func (d *day12) Part1() string {
	caves, small, start, end := d.getBitStart()
	queue := make([]node, 1, 16)
	queue[0] = node{
		path:    start,
		current: start,
	}
	ct := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, c := range caves[n.current] {
			if c == end {
				ct++
				continue
			}
			if n.path&small&c == c {
				continue
			}
			queue = append(queue, node{
				path:    n.path | c,
				current: c,
			})
		}
	}
	return fmt.Sprint(ct)
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

func (d *day12) getBitStart() (caves map[int][]int, small, start, end int) {
	caves = make(map[int][]int, 16)
	idMap := make(map[string]int, 16)
	small = 0
	nextId := 1
	getId := func(c string) int {
		if id, ok := idMap[c]; ok {
			return id
		}
		idMap[c] = nextId
		if c[0] >= 'a' {
			small |= nextId
		}
		nextId = nextId << 1
		return idMap[c]
	}
	getCave := func(c string) int {
		id := getId(c)
		if _, ok := caves[id]; !ok {
			caves[id] = make([]int, 0, 16)
		}
		return id
	}
	start = getCave("start")
	end = getCave("end")
	ensurePath := func(a, b int) {
		if b == start {
			return
		}
		idx := sort.SearchInts(caves[a], b)
		if len(caves[a]) == idx {
			caves[a] = append(caves[a], b)
			return
		}
		if caves[a][idx] != b {
			caves[a] = append(caves[a][:idx+1], caves[a][idx:]...)
			caves[a][idx] = b
		}
	}
	for _, l := range d.input {
		parts := strings.Split(l, "-")
		left, right := getCave(parts[0]), getCave(parts[1])
		ensurePath(left, right)
		ensurePath(right, left)
	}
	return
}

func New() runner.Day {
	return &day12{}
}
