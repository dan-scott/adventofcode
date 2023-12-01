package day12

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"sort"
	"strings"
)

const (
	START = 1
	END   = 2
)

type day12 struct {
	input []string
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
	caves, small := d.getBitStart()
	queue := make([]node, 1, 16)
	queue[0] = node{
		path:    START,
		current: START,
	}
	ct := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, c := range caves[n.current] {
			if c == END {
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

func (d *day12) Part2() string {
	caves, small := d.getBitStart()
	queue := make([]node, 1, 16)
	queue[0] = node{
		path:    START,
		current: START,
		doubled: false,
	}
	ct := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, c := range caves[n.current] {
			if c == END {
				ct++
				continue
			}
			doubled := n.doubled
			if n.path&small&c == c {
				if doubled {
					continue
				} else {
					doubled = true
				}
			}
			queue = append(queue, node{
				path:    n.path | c,
				current: c,
				doubled: doubled,
			})
		}
	}
	return fmt.Sprint(ct)
}

func (d *day12) getBitStart() (caves map[int][]int, small int) {
	caves = make(map[int][]int, 16)
	caves[START] = make([]int, 0, 16)
	caves[END] = make([]int, 0, 16)
	idMap := make(map[string]int, 16)
	idMap["start"] = START
	idMap["end"] = END
	small = 0
	nextId := END << 1
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
	ensurePath := func(a, b int) {
		if b == START {
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

func New() runner.LegacyDay {
	return &day12{}
}
