package day14

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/inputs"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"math"
	"sort"
	"strconv"
	"strings"
)

type day14 struct {
	blocks []block
}

type block struct {
	mask       string
	maskFloats []uint64
	ops        []op
}

type op struct {
	addr uint64
	val  uint64
}

func maskVal(mask string, in uint64) uint64 {
	bVal := fmt.Sprintf("%036b", in)
	if len(bVal) > 36 {
		bVal = bVal[len(bVal)-36:]
	}
	val := uint64(0)
	for i := 35; i >= 0; i-- {
		if mask[i] == '1' || (mask[i] == 'X' && bVal[i] == '1') {
			val += 1 << (35 - i)
		}
	}
	return val
}

func maskAddr(mask string, maskFloats []uint64, in uint64) []uint64 {
	bVal := fmt.Sprintf("%036b", in)
	if len(bVal) > 36 {
		bVal = bVal[len(bVal)-36:]
	}
	val := uint64(0)
	for i := 35; i >= 0; i-- {
		if mask[i] == '1' || (mask[i] == '0' && bVal[i] == '1') {
			val += 1 << (35 - i)
		}
	}
	vals := make([]uint64, len(maskFloats))
	for i, m := range maskFloats {
		vals[i] = m + val
	}

	return vals
}

func (d *day14) Open() {
	d.loadLines(inputs.LinesAsString(14))
}

func (d *day14) Close() {
	d.blocks = nil
}

func (d *day14) loadLines(lines []string) {
	blocks := make([]block, 0)
	b := blockForMask(lines[0])
	for _, l := range lines[1:] {
		if strings.Contains(l, "mask = ") {
			blocks = append(blocks, b)
			b = blockForMask(l)
		} else {
			op := parseOp(l)
			b.ops = append(b.ops, op)
		}
	}
	blocks = append(blocks, b)
	d.blocks = blocks
}

func blockForMask(l string) block {
	mask := l[7:]
	floats := make([]int, 0)
	for i, c := range mask {
		if c == 'X' {
			floats = append(floats, 35-i)
		}
	}
	sort.Slice(floats, func(i int, j int) bool { return i > j })
	maskFloats := make([]uint64, int(math.Pow(2, float64(len(floats)))))
	// pre-calculate all the float values that would be added to a masked address
	for i, r := range floats {
		b := true
		m := int(math.Pow(float64(2), float64(i)))
		for j := range maskFloats {
			if j%m == 0 {
				b = !b
			}
			if b {
				maskFloats[j] += 1 << r
			}
		}
	}

	ops := make([]op, 0)
	return block{mask, maskFloats, ops}
}

func parseOp(l string) op {
	parts := strings.Split(l, "] = ")
	mem, _ := strconv.ParseUint(parts[0][4:], 10, 64)
	val, _ := strconv.ParseUint(parts[1], 10, 64)
	return op{addr: mem, val: val}
}

func (d *day14) Part1() string {
	mem := make(map[uint64]uint64)
	for _, b := range d.blocks {
		for _, o := range b.ops {
			mem[o.addr] = maskVal(b.mask, o.val)
		}
	}
	total := uint64(0)
	for _, v := range mem {
		total += v
	}

	return fmt.Sprint(total)
}

func (d *day14) Part2() string {
	mem := make(map[uint64]uint64)
	for _, b := range d.blocks {
		for _, o := range b.ops {
			for _, a := range maskAddr(b.mask, b.maskFloats, o.addr) {
				mem[a] = o.val
			}
		}
	}
	total := uint64(0)
	for _, v := range mem {
		total += v
	}
	return fmt.Sprint(total)
}

func New() internal.Day {
	return &day14{}
}
