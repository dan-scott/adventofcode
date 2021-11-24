package day08

import (
	"errors"
	"fmt"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/inputs"
	"gitlab.com/danscott/adventofcode/adventofcode_2020_go/internal"
	"strconv"
)

var (
	errLoop = errors.New("loop detected")
)

type day08 struct {
	ops []op
}

type vm struct {
	ops     []op
	pc      int
	ac      int
	history map[int]interface{}
}

func vmOf(ops []op) *vm {
	return &vm{
		ops:     ops,
		history: make(map[int]interface{}),
	}
}

func (v *vm) run() error {
	for v.pc < len(v.ops) {
		if err := v.tick(); err != nil {
			return err
		}
	}
	return nil
}

func (v *vm) tick() error {
	op := v.ops[v.pc]
	switch op.op {
	case "acc":
		v.ac += op.val
		v.pc++
		break
	case "jmp":
		v.pc += op.val
		break
	default:
		v.pc++
	}
	if _, ok := v.history[v.pc]; ok {
		return errLoop
	}
	v.history[v.pc] = nil
	return nil
}

type op struct {
	op  string
	val int
}

func (d *day08) Open() {
	lines := inputs.LinesAsString(8)
	d.loadLines(lines)
}

func (d *day08) loadLines(lines []string) {
	d.ops = make([]op, len(lines))
	for i, l := range lines {
		val, _ := strconv.ParseInt(l[4:], 10, 32)
		d.ops[i] = op{
			op:  l[:3],
			val: int(val),
		}
	}
}

func (d *day08) Close() {
	d.ops = nil
}

func (d *day08) Part1() string {
	vm := vmOf(d.ops)
	_ = vm.run()
	return fmt.Sprint(vm.ac)
}

func (d *day08) Part2() string {
	for i, o := range d.ops {
		if o.op == "jmp" {
			o2 := make([]op, len(d.ops))
			copy(o2, d.ops)
			o2[i] = op{
				op:  "nop",
				val: o.val,
			}

			vm := vmOf(o2)
			if err := vm.run(); err == nil {
				return fmt.Sprint(vm.ac)
			}
		}

		if o.op == "nop" {
			o2 := make([]op, len(d.ops))
			copy(o2, d.ops)
			o2[i] = op{
				op:  "jmp",
				val: o.val,
			}

			vm := vmOf(o2)
			if err := vm.run(); err == nil {
				return fmt.Sprint(vm.ac)
			}
		}
	}

	return "too bad"
}

func New() internal.Day {
	return &day08{}
}
