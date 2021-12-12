package day12

import "testing"

func TestDay12_Part1(t *testing.T) {
	d := &day12{}
	d.Open()
	const expected = "3369"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func TestDay12_Part2(t *testing.T) {
	d := New()
	d.Open()
	const expected = "85883"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func BenchmarkDay12_Part1(b *testing.B) {
	d := &day12{}
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay12_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
