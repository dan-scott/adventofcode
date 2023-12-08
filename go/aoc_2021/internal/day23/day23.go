package day23

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"github.com/danscott/adventofcode/go/common/vec2"
	"sort"
	"strings"
)

const (
	cX byte = '.'
	cA byte = 'A'
	cB byte = 'B'
	cC byte = 'C'
	cD byte = 'D'
)

var order = []byte{cA, cB, cC, cD}

func supports(x int) byte {
	return order[x/2-1]
}

var costMap = map[byte]int{
	cA: 1,
	cB: 10,
	cC: 100,
	cD: 1000,
}

var targetHall = map[byte]int{
	cA: 2,
	cB: 4,
	cC: 6,
	cD: 8,
}

type crab struct {
	pos      vec2.Vec2
	cost     int
	t        byte
	hasMoved bool
}

func (c crab) String() string {
	return fmt.Sprintf("[%c:%s:%v]", c.t, c.pos, c.hasMoved)
}

func crb(x, y int, t byte) crab {
	return crab{
		pos:      vec2.Of(x, y),
		cost:     costMap[t],
		t:        t,
		hasMoved: false,
	}
}

type burrow struct {
	hallSpots map[vec2.Vec2]interface{}
	roomSpots map[vec2.Vec2]interface{}
	crabs     []crab
	cmap      map[vec2.Vec2]crab
}

func signature(crabs []crab) string {
	crabIds := make([]string, len(crabs))
	for i, c := range crabs {
		crabIds[i] = c.String()
	}
	sort.Strings(crabIds)
	return strings.Join(crabIds, " ")
}

func (b *burrow) indexCrabs() {
	b.cmap = make(map[vec2.Vec2]crab, len(b.crabs))
	for _, c := range b.crabs {
		b.cmap[c.pos] = c
	}
}

func (b *burrow) winningSignature() string {
	crabs := make([]crab, 0, len(b.roomSpots))
	for pos := range b.roomSpots {
		c := crb(pos.X, pos.Y, supports(pos.X))
		c.hasMoved = true
		crabs = append(crabs, c)
	}
	return signature(crabs)
}

func (b *burrow) hallCrabs() []crab {
	hcs := make([]crab, 0, len(b.crabs))
	for _, c := range b.crabs {
		if c.pos.Y == 0 {
			if c.pos.X == 0 {
				if _, ok := b.cmap[vec2.Of(1, 0)]; ok {
					continue
				}
			}
			if c.pos.X == 10 {
				if _, ok := b.cmap[vec2.Of(9, 0)]; ok {
					continue
				}
			}
			hcs = append(hcs, c)
		}
	}
	sort.SliceStable(hcs, func(i, j int) bool {
		return hcs[i].cost < hcs[j].cost
	})
	return hcs
}

func (b *burrow) topRoomCrabs() []crab {
	var crabs []crab
	if c, ok := b.topRoomCrab(2); ok {
		crabs = append(crabs, *c)
	}
	if c, ok := b.topRoomCrab(4); ok {
		crabs = append(crabs, *c)
	}
	if c, ok := b.topRoomCrab(6); ok {
		crabs = append(crabs, *c)
	}
	if c, ok := b.topRoomCrab(8); ok {
		crabs = append(crabs, *c)
	}
	sort.SliceStable(crabs, func(i, j int) bool {
		return crabs[i].cost < crabs[j].cost
	})
	return crabs
}

func (b *burrow) topRoomCrab(x int) (*crab, bool) {
	var c *crab
	for _, cr := range b.crabs {
		if cr.pos.X == x && !cr.hasMoved && (c == nil || c.pos.Y < cr.pos.Y) {
			c = &cr
		}
	}
	return c, c != nil
}

func (b *burrow) canEnterRoom(c crab) (vec2.Vec2, bool) {
	x := targetHall[c.t]
	if tc, has := b.topRoomCrab(targetHall[c.t]); has {
		if tc.t == c.t && tc.hasMoved {
			return tc.pos.Add(vec2.Of(x, -1)), true
		} else {
			return vec2.Of(0, 0), false
		}
	}
	return vec2.Of(x, 2), true
}

func (b *burrow) solve() int {
	win := b.winningSignature()
	cost := 0
SolveLoop:
	for signature(b.crabs) != win {
		hcs := b.hallCrabs()
		rcs := b.topRoomCrabs()
		for len(hcs) > 0 && len(rcs) > 0 {
			var n crab
			if len(hcs) == 0 {
				n = rcs[0]
				rcs = rcs[1:]
			} else if len(rcs) == 0 || hcs[0].cost < rcs[0].cost {
				n = hcs[0]
				hcs = hcs[1:]
			} else {
				n = rcs[0]
				rcs = rcs[1:]
			}

			if n.pos.Y == 0 {
				if dest, can := b.canEnterRoom(n); can {
					for _, c := range b.crabs {
						if c.pos.Y == 0 && ((c.pos.X > n.pos.X && c.pos.X < dest.X) || (c.pos.X < n.pos.X && c.pos.X > dest.X)) {
							continue SolveLoop
						}
					}
					cost += b.moveCrab(n, dest)
					continue SolveLoop
				}
			} else {
				if dest, can := b.findFarthestHall(n); can {
					cost += b.moveCrab(n, dest)
					continue SolveLoop
				}
			}
		}
		panic("no moves!")
	}

	return cost
}

func (b *burrow) moveCrab(c crab, dest vec2.Vec2) int {
	cost := c.pos.MhdTo(dest)
	c.pos = dest
	c.hasMoved = true
	b.indexCrabs()
	return cost * c.cost
}

func (b *burrow) findFarthestHall(c crab) (vec2.Vec2, bool) {
	foundL := false
	foundR := false
	lh := vec2.Of(-1, -1)
	rh := lh
	for x := c.pos.X; x >= 0; x-- {
		pt := vec2.Of(x, 0)
		_, isHallSpot := b.hallSpots[pt]
		_, isOccupied := b.cmap[pt]
		if isOccupied {
			break
		}
		if isHallSpot {
			foundL = true
			lh = pt
		}
	}
	for x := c.pos.X; x <= 10; x++ {
		pt := vec2.Of(x, 0)
		_, isHallSpot := b.hallSpots[pt]
		_, isOccupied := b.cmap[pt]
		if isOccupied {
			break
		}
		if isHallSpot {
			foundR = true
			rh = pt
		}
	}
	if !foundR && !foundL {
		return lh, false
	}
	if foundL && foundR {
		if c.pos.MhdTo(lh) > c.pos.MhdTo(rh) {
			return lh, true
		} else {
			return rh, true
		}
	}
	if foundL {
		return lh, true
	}
	return rh, true
}

type day23 struct {
	input []string
}

func (d *day23) Open() {
	d.input = inputs.LinesAsString(2021, 23)
}

func (d *day23) Close() {
	d.input = nil
}

func (d *day23) Part1() string {
	b := parseInput(d.input)
	cost := b.solve()
	return fmt.Sprint(cost)
}

func (d *day23) Part2() string {
	return ""
}

func parseInput(input []string) *burrow {
	b := &burrow{
		hallSpots: map[vec2.Vec2]interface{}{
			vec2.Of(0, 0):  nil,
			vec2.Of(1, 0):  nil,
			vec2.Of(3, 0):  nil,
			vec2.Of(5, 0):  nil,
			vec2.Of(7, 0):  nil,
			vec2.Of(9, 0):  nil,
			vec2.Of(10, 0): nil,
		},
		roomSpots: map[vec2.Vec2]interface{}{
			vec2.Of(2, 1): nil,
			vec2.Of(2, 2): nil,
			vec2.Of(4, 1): nil,
			vec2.Of(4, 2): nil,
			vec2.Of(6, 1): nil,
			vec2.Of(6, 2): nil,
			vec2.Of(8, 1): nil,
			vec2.Of(8, 2): nil,
		},
		crabs: []crab{
			crb(2, 1, input[2][3]),
			crb(2, 2, input[3][3]),
			crb(4, 1, input[2][5]),
			crb(4, 2, input[3][5]),
			crb(6, 1, input[2][7]),
			crb(6, 2, input[3][7]),
			crb(8, 1, input[2][9]),
			crb(8, 2, input[3][9]),
		},
	}
	b.indexCrabs()
	for _, c := range b.crabs {
		if c.t == supports(c.pos.X) && c.pos.Y == 2 {
			c.hasMoved = true
		}
	}
	return b
}

func New() runner.LegacyDay {
	return &day23{}
}
