package day03

import (
	"fmt"
	"github.com/danscott/adventofcode/go/common/inputs"
	"github.com/danscott/adventofcode/go/common/runner"
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
	zct := zeroCt(d.bins)
	bitCt := len(d.bins[0])
	gamma := 0
	epsilon := 0
	half := len(d.bins) / 2
	for i := 0; i < bitCt; i++ {
		val := 1 << (bitCt - i - 1)
		if zct[i] >= half {
			gamma |= val
		} else {
			epsilon |= val
		}
	}

	return fmt.Sprint(gamma * epsilon)
}

func zeroCt(lines []string) []int {
	zct := make([]int, len(lines[0]))
	for _, n := range lines {
		for i, c := range n {
			if c == '0' {
				zct[i]++
			}
		}
	}
	return zct
}

func (d *day03) Part2() string {
	ogChan := make(chan int64, 1)
	go func(filter []string, ch chan int64) {
		bit := 0
		zct := zeroCt(filter)
		for len(filter) > 1 {
			common := uint8('1')
			if zct[bit] > len(filter)/2 {
				common = '0'
			}
			nOgr := make([]string, 0)
			for _, b := range filter {
				if b[bit] == common {
					nOgr = append(nOgr, b)
				} else {
					for i, n := range b {
						if n == '0' {
							zct[i]--
						}
					}
				}
			}
			filter = nOgr
			bit++
		}
		ogr, _ := strconv.ParseInt(filter[0], 2, 32)
		ch <- ogr
	}(d.bins, ogChan)

	co2Chan := make(chan int64, 1)
	go func(filter []string, ch chan int64) {
		bit := 0
		zct := zeroCt(filter)
		for len(filter) > 1 {
			common := uint8('0')
			if zct[bit] > len(filter)/2 {
				common = '1'
			}
			nOgr := make([]string, 0)
			for _, b := range filter {
				if b[bit] == common {
					nOgr = append(nOgr, b)
				} else {
					for i, n := range b {
						if n == '0' {
							zct[i]--
						}
					}
				}
			}
			filter = nOgr
			bit++
		}
		co2, _ := strconv.ParseInt(filter[0], 2, 32)
		ch <- co2
	}(d.bins, co2Chan)

	ogr := <-ogChan
	co2 := <-co2Chan
	return fmt.Sprint(ogr * co2)
}

func New() runner.LegacyDay {
	return &day03{}
}
