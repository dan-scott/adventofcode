package day14

import "testing"

func TestMaskVals(t *testing.T) {
	lines := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	d := &day14{}
	d.loadLines(lines)
	if d.Part1() != "165" {
		t.Fatalf("Expected 165 but got %s", d.Part1())
	}
}

func TestMaskAddrs(t *testing.T) {
	lines := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	d := &day14{}
	d.loadLines(lines)
	if d.Part2() != "208" {
		t.Fatalf("Expected 208 but got %s", d.Part2())
	}
}
