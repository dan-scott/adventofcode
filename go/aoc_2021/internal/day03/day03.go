package day03

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"strconv"
)

type day03 struct {
	bins []string
}

func (d *day03) Open() {
	d.loadLines(inputs.LinesAsString(2021, 3))
}

func (d *day03) loadLines(lines []string) {
	d.bins = lines
}

func (d *day03) Close() {
	d.bins = nil
}

func (d *day03) Part1() string {
	bitCt, zct, oct := bitCts(d.bins)
	gamma := 0
	epsilon := 0
	for i := 0; i < bitCt; i++ {
		val := 1 << (bitCt - i - 1)
		if zct[i] > oct[i] {
			gamma += val
		} else {
			epsilon += val
		}
	}

	return fmt.Sprint(gamma * epsilon)
}

func bitCts(lines []string) (bitCt int, zct, oct []int) {
	bitCt = len(lines[0])
	zct = make([]int, bitCt)
	oct = make([]int, bitCt)
	for _, n := range lines {
		for i, c := range n {
			if c == '0' {
				zct[i]++
			} else {
				oct[i]++
			}
		}
	}
	return
}

func (d *day03) Part2() string {
	filter := make([]string, len(d.bins))
	copy(filter, d.bins)
	bit := 0
	for len(filter) > 1 {
		_, zct, oct := bitCts(filter)
		common := uint8('1')
		if zct[bit] > oct[bit] {
			common = '0'
		}
		nOgr := make([]string, 0)
		for _, b := range filter {
			if b[bit] == common {
				nOgr = append(nOgr, b)
			}
		}
		filter = nOgr
		bit++
	}
	ogr, _ := strconv.ParseInt(filter[0], 2, 32)
	filter = make([]string, len(d.bins))
	copy(filter, d.bins)
	bit = 0
	for len(filter) > 1 {
		_, zct, oct := bitCts(filter)
		common := uint8('0')
		if zct[bit] > oct[bit] {
			common = '1'
		}
		nOgr := make([]string, 0)
		for _, b := range filter {
			if b[bit] == common {
				nOgr = append(nOgr, b)
			}
		}
		filter = nOgr
		bit++
	}
	co2, _ := strconv.ParseInt(filter[0], 2, 32)
	return fmt.Sprint(ogr * co2)
}

func New() runner.Day {
	return &day03{}
}
