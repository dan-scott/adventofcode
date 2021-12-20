package day18

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"math"
)

type day18 struct {
	lines []string
}

type node interface {
	asRegular() (*regular, bool)
	asPair() (*pair, bool)
	magnitude() int
	nextExplode(level int) (*pair, bool)
	nextSplittable() (*regular, bool)
	regularDfs(regs []*regular) []*regular
}

type pair struct {
	left   node
	right  node
	parent *pair
}

func (p *pair) String() string {
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func (p *pair) regularDfs(regs []*regular) []*regular {
	regs = p.left.regularDfs(regs)
	regs = p.right.regularDfs(regs)
	return regs
}

func (p *pair) add(o *pair) *pair {
	np := &pair{
		left:  p,
		right: o,
	}
	p.parent = np
	o.parent = np
	return np.reduce()
}

func (p *pair) reduce() *pair {
	for {
		if p.tryExplode() {
			continue
		}
		if !p.trySplit() {
			break
		}
	}
	return p
}

func (p *pair) tryExplode() bool {
	if p.parent != nil {
		return p.parent.tryExplode()
	}
	explode, has := p.nextExplode(0)
	if !has {
		return false
	}
	var regulars = p.regularDfs(make([]*regular, 0))

	idx := 0
	for i, r := range regulars {
		if r == explode.left {
			idx = i
			break
		}
	}
	if idx > 0 {
		leftVal, _ := explode.left.asRegular()
		regulars[idx-1].value += leftVal.value
	}
	if idx+1 < len(regulars)-1 {
		rightVal, _ := explode.right.asRegular()
		regulars[idx+2].value += rightVal.value
	}
	exParent := explode.parent
	z := &regular{value: 0, parent: exParent}
	if exParent.left == explode {
		exParent.left = z
	} else {
		exParent.right = z
	}
	explode.parent = nil
	return true
}

func (p *pair) trySplit() bool {
	if p.parent != nil {
		return p.parent.trySplit()
	}
	splittable, has := p.nextSplittable()
	if !has {
		return false
	}
	m := float64(splittable.value) / 2
	l := math.Floor(m)
	r := math.Ceil(m)
	np := &pair{parent: splittable.parent}
	lv := &regular{value: int(l), parent: np}
	rv := &regular{value: int(r), parent: np}
	np.left = lv
	np.right = rv
	if splittable.parent.left == splittable {
		splittable.parent.left = np
	} else {
		splittable.parent.right = np
	}
	splittable.parent = nil
	return true
}

func (p *pair) nextExplode(level int) (*pair, bool) {
	e, found := p.left.nextExplode(level + 1)

	if !found {
		e, found = p.right.nextExplode(level + 1)
	}

	if !found && level >= 4 {
		e, found = p, true
	}

	return e, found
}

func (p *pair) nextSplittable() (*regular, bool) {
	r, found := p.left.nextSplittable()
	if !found {
		r, found = p.right.nextSplittable()
	}
	return r, found
}

func (p *pair) asRegular() (*regular, bool) {
	return nil, false
}

func (p *pair) asPair() (*pair, bool) {
	return p, true
}

func (p *pair) magnitude() int {
	return (3 * p.left.magnitude()) + (2 * p.right.magnitude())
}

type regular struct {
	value  int
	parent *pair
}

func (r *regular) String() string {
	return fmt.Sprint(r.value)
}

func (r *regular) regularDfs(regs []*regular) []*regular {
	regs = append(regs, r)
	return regs
}

func (r *regular) nextExplode(_ int) (*pair, bool) {
	return nil, false
}

func (r *regular) nextSplittable() (*regular, bool) {
	if r.value > 9 {
		return r, true
	}
	return nil, false
}

func (r *regular) asRegular() (*regular, bool) {
	return r, true
}

func (r *regular) asPair() (*pair, bool) {
	return nil, false
}

func (r *regular) magnitude() int {
	return r.value
}

func (d *day18) Open() {
	d.lines = inputs.LinesAsString(2021, 18)
}

func (d *day18) Close() {
	d.lines = nil
}

func (d *day18) Part1() string {
	pairs := d.parseInput()
	root := pairs[0]
	for _, next := range pairs[1:] {
		root = root.add(next)
	}
	return fmt.Sprint(root.magnitude())
}

func (d *day18) Part2() string {
	max := 0
	for a := 0; a < len(d.lines); a++ {
		for b := 0; b < len(d.lines); b++ {
			if a != b {
				ap := parseLine(d.lines[a])
				bp := parseLine(d.lines[b])
				ans := ap.add(bp).magnitude()
				if ans > max {
					max = ans
				}
			}
		}
	}
	return fmt.Sprint(max)
}

func (d *day18) parseInput() []*pair {
	pairs := make([]*pair, len(d.lines))
	for i, line := range d.lines {
		pairs[i] = parseLine(line)
	}
	return pairs
}

type stack struct {
	pairs []*pair
}

func newStack() *stack {
	return &stack{pairs: make([]*pair, 0, 4)}
}

func (s *stack) pop() (*pair, bool) {
	if len(s.pairs) == 0 {
		return nil, false
	}
	p := s.pairs[len(s.pairs)-1]
	s.pairs = s.pairs[:len(s.pairs)-1]
	return p, true
}

func (s *stack) push(p *pair) {
	s.pairs = append(s.pairs, p)
}

func parseLine(line string) *pair {
	s := newStack()
	current := &pair{}
	for _, c := range line[1:] {
		if c == '[' {
			newPair := &pair{parent: current}
			if current.left == nil {
				current.left = newPair
			} else {
				current.right = newPair
			}
			s.push(current)
			current = newPair
		} else if c == ',' {
			continue
		} else if c == ']' {
			if p, ok := s.pop(); ok {
				current = p
			} else {
				return current
			}
		} else if current.left == nil {
			current.left = &regular{
				value:  int(c - '0'),
				parent: current,
			}
		} else {
			current.right = &regular{
				value:  int(c - '0'),
				parent: current,
			}
		}
	}
	return current
}

func New() runner.Day {
	return &day18{}
}
