package day07

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/aoc_2020/internal"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"strings"
)

type day07 struct {
	bags map[bagName]*bag
}

type bagName = string
type bag struct {
	name     bagName
	parents  []*bagLink
	children []*bagLink
}
type bagLink struct {
	parent     *bag
	child      *bag
	childCount uint8
}

func newBag(name bagName) *bag {
	return &bag{
		name:     name,
		parents:  make([]*bagLink, 0),
		children: make([]*bagLink, 0),
	}
}

func (b *bag) isRoot() bool {
	return len(b.parents) == 0
}

func (b *bag) isLeaf() bool {
	return len(b.children) == 0
}

func (b *bag) addChild(l *bagLink) {
	b.children = append(b.children, l)
}

func (b *bag) addParent(l *bagLink) {
	b.parents = append(b.parents, l)
}

func (b *bag) countDistinctParents() uint {
	n := make(map[bagName]interface{}, 0)
	b.countDistinctParentsWithMap(n)
	return uint(len(n) - 1)
}

func (b *bag) countDistinctParentsWithMap(n map[bagName]interface{}) {
	if _, ok := n[b.name]; ok {
		return
	} else {
		n[b.name] = nil
	}
	for _, p := range b.parents {
		p.parent.countDistinctParentsWithMap(n)
	}
}

func (b *bag) countKids() uint {
	total := uint(1)
	for _, c := range b.children {
		total += uint(c.childCount) * c.child.countKids()
	}
	return total
}

func (d *day07) Open() {
	lines := inputs.LinesAsString(2020, 7)
	d.loadLines(lines)
}

func (d *day07) Close() {
	d.bags = nil
}

func (d *day07) Part1() string {
	return fmt.Sprint(d.bags["shiny gold"].countDistinctParents())
}

func (d *day07) Part2() string {
	return fmt.Sprint(d.bags["shiny gold"].countKids() - 1)
}

func (d *day07) loadLines(lines []string) {
	d.bags = make(map[bagName]*bag)
	for _, l := range lines {
		bn, children := parseLine(l)
		if _, ok := d.bags[bn]; !ok {
			d.bags[bn] = newBag(bn)
		}
		for _, c := range children {
			if _, ok := d.bags[c.name]; !ok {
				d.bags[c.name] = newBag(c.name)
			}
			link := &bagLink{
				parent:     d.bags[bn],
				child:      d.bags[c.name],
				childCount: c.count,
			}
			d.bags[bn].addChild(link)
			d.bags[c.name].addParent(link)
		}
	}
}

type child struct {
	name  bagName
	count uint8
}

type line struct {
	parts []string
}

func (l *line) chompName() bagName {
	parts := l.chomp(2)
	return fmt.Sprintf("%s %s", parts[0], parts[1])
}

func (l *line) chompCount() uint8 {
	cs := l.chomp(1)
	return uint8(internal.MustParseUint(cs[0]))
}

func (l *line) chomp(words uint) []string {
	p := l.parts[:words]
	l.parts = l.parts[words:]
	return p
}

func parseLine(l string) (bn bagName, children []child) {
	line := &line{parts: strings.Split(l, " ")}
	bn = line.chompName()
	line.chomp(2) // chomp bags contain
	children = make([]child, 0)

	for len(line.parts) > 0 {
		ct := line.chompCount()
		name := line.chompName()
		if name == "other bags." {
			return
		}
		line.chomp(1)
		children = append(children, child{
			name:  name,
			count: ct,
		})
	}
	return
}

func New() runner.Day {
	return &day07{}
}
