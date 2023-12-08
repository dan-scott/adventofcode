package day10

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
	"sort"
)

var (
	corruptScore    map[byte]int
	incompleteScore map[byte]int
	close           map[byte]byte
	open            map[byte]byte
)

func init() {
	corruptScore = make(map[byte]int, 4)
	corruptScore[')'] = 3
	corruptScore[']'] = 57
	corruptScore['}'] = 1197
	corruptScore['>'] = 25137

	incompleteScore = make(map[byte]int, 4)
	incompleteScore['('] = 1
	incompleteScore['['] = 2
	incompleteScore['{'] = 3
	incompleteScore['<'] = 4

	close = make(map[byte]byte, 4)
	close[')'] = '('
	close[']'] = '['
	close['}'] = '{'
	close['>'] = '<'

	open = make(map[byte]byte, 4)
	open['('] = ')'
	open['['] = ']'
	open['{'] = '}'
	open['<'] = '>'
}

type day10 struct {
	lines []string
}

func (d *day10) Open() {
	d.loadLines(inputs.LinesAsString(2021, 10))
}

func (d *day10) Close() {
	d.lines = nil
}

func (d *day10) Part1() string {
	total := 0
	for _, l := range d.lines {
		if score, corrupted := getScore(l); corrupted {
			total += score
		}
	}
	return fmt.Sprint(total)
}

func getScore(line string) (score int, isCorrupted bool) {
	stack := make([]byte, 0, len(line))
	stack = append(stack, line[0])
	for i := 1; i < len(line); i++ {
		if _, ok := open[line[i]]; ok {
			stack = append(stack, line[i])
		} else if close[line[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return corruptScore[line[i]], true
		}
	}
	total := 0
	for i := len(stack) - 1; i >= 0; i-- {
		total = (total * 5) + incompleteScore[stack[i]]
	}
	return total, false
}

func (d *day10) Part2() string {
	scores := make([]int, 0, len(d.lines))
	for _, l := range d.lines {
		if score, corrupted := getScore(l); !corrupted {
			scores = append(scores, score)
		}
	}
	sort.SliceStable(scores, func(i int, j int) bool {
		return scores[i] < scores[j]
	})
	return fmt.Sprint(scores[len(scores)/2])
}

func (d *day10) loadLines(lines []string) {
	d.lines = lines
}

func New() runner.LegacyDay {
	return &day10{}
}
