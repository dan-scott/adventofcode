package day16

import (
	"fmt"
	"gitlab.com/danscott/adventofcode/go/common/inputs"
	"gitlab.com/danscott/adventofcode/go/common/runner"
	"math"
	"strconv"
)

type packetType = int

const (
	tSUM packetType = iota
	tPROD
	tMIN
	tMAX
	tLIT
	tGT
	tLT
	tEQ
)

var hexMap = make(map[rune]string)

func init() {
	hexMap['0'] = "0000"
	hexMap['1'] = "0001"
	hexMap['2'] = "0010"
	hexMap['3'] = "0011"
	hexMap['4'] = "0100"
	hexMap['5'] = "0101"
	hexMap['6'] = "0110"
	hexMap['7'] = "0111"
	hexMap['8'] = "1000"
	hexMap['9'] = "1001"
	hexMap['A'] = "1010"
	hexMap['B'] = "1011"
	hexMap['C'] = "1100"
	hexMap['D'] = "1101"
	hexMap['E'] = "1110"
	hexMap['F'] = "1111"
}

type day16 struct {
	hex string
}

type parser struct {
	i    int
	bits string
}

func (p *parser) chompInt(bitSize int) int {
	v, _ := strconv.ParseInt(p.chompBits(bitSize), 2, 64)
	return int(v)
}

func (p *parser) chompBits(bitSize int) string {
	s := p.bits[p.i : p.i+bitSize]
	p.i += bitSize
	return s
}

func (p *parser) chompVersion() int {
	return p.chompInt(3)
}

func (p *parser) chompTypeId() int {
	return p.chompInt(3)
}

func (p *parser) chompPacket() (vSum, val int) {
	vSum = p.chompVersion()
	t := p.chompTypeId()
	if t == tLIT {
		val = p.chompLiteral()
	} else {
		c, v := p.chompOperator(t)
		vSum += c
		val = v
	}
	return
}

func (p *parser) chompLiteral() int {
	readingLit := true
	litBitCt := 6
	litStr := ""
	for readingLit {
		litBitCt += 5
		leader := p.chompBits(1)
		litStr = fmt.Sprintf("%s%s", litStr, p.chompBits(4))
		if leader == "0" {
			readingLit = false
		}
	}
	v, _ := strconv.ParseInt(litStr, 2, 64)
	return int(v)
}

func (p *parser) chompOperator(t packetType) (vSum, val int) {
	vSum = 0
	val = 0
	var values []int
	if p.chompBits(1) == "0" {
		vSum, values = p.chompLengthOp()
	} else {
		vSum, values = p.chompNumOp()
	}
	switch t {
	case tSUM:
		val = sum(values)
	case tPROD:
		val = prod(values)
	case tMIN:
		val = min(values)
	case tMAX:
		val = max(values)
	case tGT:
		if values[0] > values[1] {
			val = 1
		}
	case tLT:
		if values[0] < values[1] {
			val = 1
		}
	case tEQ:
		if values[0] == values[1] {
			val = 1
		}
	}
	return
}

func max(values []int) int {
	m := 0
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}

func min(values []int) int {
	m := math.MaxInt
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m
}

func prod(values []int) int {
	p := 1
	for _, v := range values {
		p *= v
	}
	return p
}

func sum(values []int) int {
	s := 0
	for _, v := range values {
		s += v
	}
	return s
}

func (p *parser) chompLengthOp() (vSum int, values []int) {
	length := p.chompInt(15)
	start := p.i
	vSum = 0
	values = make([]int, 0)
	for (p.i - start) < length {
		s, v := p.chompPacket()
		vSum += s
		values = append(values, v)
	}
	return
}

func (p *parser) chompNumOp() (vSum int, values []int) {
	num := p.chompInt(11)
	vSum = 0
	values = make([]int, 0, num)
	for i := 0; i < num; i++ {
		s, v := p.chompPacket()
		vSum += s
		values = append(values, v)
	}
	return
}

func (d *day16) Open() {
	d.hex = inputs.LinesAsString(2021, 16)[0]
}

func (d *day16) Close() {
}

func (d *day16) Part1() string {
	p := d.getParser()
	vSum, _ := p.chompPacket()
	return fmt.Sprint(vSum)
}

func (d *day16) Part2() string {
	p := d.getParser()
	_, val := p.chompPacket()
	return fmt.Sprint(val)
}

func (d *day16) getParser() *parser {
	bits := ""
	for _, c := range d.hex {
		bits = fmt.Sprintf("%s%s", bits, hexMap[c])
	}
	return &parser{bits: bits}
}

func New() runner.LegacyDay {
	return &day16{}
}
