package day15

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		line     string
		expected string
	}{
		{"0,3,6", "436"},
		{"1,3,2", "1"},
		{"2,1,3", "10"},
		{"1,2,3", "27"},
		{"2,3,1", "78"},
		{"3,2,1", "438"},
		{"3,1,2", "1836"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Init %s", tc.line), func(t *testing.T) {
			d := &day15{}
			d.load(tc.line)
			actual := d.Part1()
			if actual != tc.expected {
				t.Fatalf("Expected %s but got %s", tc.expected, actual)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		line     string
		expected string
	}{
		{"0,3,6", "175594"},
		{"1,3,2", "2578"},
		{"2,1,3", "3544142"},
		{"1,2,3", "261214"},
		{"2,3,1", "6895259"},
		{"3,2,1", "18"},
		{"3,1,2", "362"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Init %s", tc.line), func(t *testing.T) {
			d := &day15{}
			d.load(tc.line)
			actual := d.Part2()
			if actual != tc.expected {
				t.Fatalf("Expected %s but got %s", tc.expected, actual)
			}
		})
	}
}

func BenchmarkPart2(b *testing.B) {
	d := &day15{}
	d.load("1,0,16,5,17,4")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}
