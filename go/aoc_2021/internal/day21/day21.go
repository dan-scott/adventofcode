package day21

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"math"
	"strconv"
	"strings"
)

type day21 struct {
	inputs []string
}

func (d *day21) Open() {
	d.inputs = inputs.LinesAsString(2021, 21)
}

func (d *day21) Close() {
	d.inputs = nil
}

func (d *day21) Part1() string {
	p1, p2 := d.parseInput()
	p1s, p2s := 0, 0
	dice := 0
	rollNext := func() int {
		r := 3*dice + 6
		dice += 3
		return r
	}
	p1turn := true
	for p1s < 1000 && p2s < 1000 {
		roll := rollNext()
		if p1turn {
			p1 = (p1 + roll) % 10
			p1s += p1 + 1
		} else {
			p2 = (p2 + roll) % 10
			p2s += p2 + 1
		}
		p1turn = !p1turn
	}
	min := int(math.Min(float64(p1s), float64(p2s)))
	return fmt.Sprint(min * dice)
}

func (d *day21) Part2() string {
	branches := genRollBranches()
	initialState := state{
		p1turn:  true,
		p1score: 0,
		p2score: 0,
	}
	initialState.p1idx, initialState.p2idx = d.parseInput()
	mem := make(map[state]outcome)
	st := &stack{[]state{}}
	st.push(initialState)
	for st.any() {
		s := st.pop()
		if _, ok := mem[s]; ok {
			continue
		}
		if s.p1score >= 21 {
			mem[s] = outcome{
				p1wins: 1,
				p2wins: 0,
			}
			continue
		}
		if s.p2score >= 21 {
			mem[s] = outcome{
				p1wins: 0,
				p2wins: 1,
			}
			continue
		}
		st.push(s)
		allSolved := true
		oc := outcome{}
		for _, roll := range branches {
			next := state{
				p1turn:  !s.p1turn,
				p1idx:   s.p1idx,
				p2idx:   s.p2idx,
				p1score: s.p1score,
				p2score: s.p2score,
			}
			if s.p1turn {
				next.p1idx = (s.p1idx + roll) % 10
				next.p1score += next.p1idx + 1
			} else {
				next.p2idx = (s.p2idx + roll) % 10
				next.p2score += next.p2idx + 1
			}
			if noc, has := mem[next]; has {
				oc.p1wins += noc.p1wins
				oc.p2wins += noc.p2wins
			} else {
				st.push(next)
				allSolved = false
			}
		}
		if allSolved {
			mem[s] = oc
		}
	}
	oc := mem[initialState]
	return fmt.Sprint(int(math.Max(float64(oc.p1wins), float64(oc.p2wins))))
}

func genRollBranches() []int {
	branches := make([]int, 0)
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				branches = append(branches, i+j+k)
			}
		}
	}
	return branches
}

func (d *day21) parseInput() (p1, p2 int) {
	p1, _ = strconv.Atoi(strings.Split(d.inputs[0], ": ")[1])
	p2, _ = strconv.Atoi(strings.Split(d.inputs[1], ": ")[1])
	p1--
	p2--
	return
}

type state struct {
	p1turn           bool
	p1idx, p2idx     int
	p1score, p2score int
}

type outcome struct {
	p1wins, p2wins int
}

type stack struct {
	states []state
}

func (s *stack) push(st state) {
	s.states = append(s.states, st)
}

func (s *stack) pop() state {
	l := len(s.states)
	t := s.states[l-1]
	s.states = s.states[:l-1]
	return t
}

func (s *stack) any() bool {
	return len(s.states) > 0
}

func New() runner.LegacyDay {
	return &day21{}
}
