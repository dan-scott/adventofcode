package day08

import "testing"

func TestDay08_Part1(t *testing.T) {
	d := &day08{}
	d.Open()
	const expected = "532"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func TestDay08_Part2(t *testing.T) {
	d := &day08{}
	d.Open()
	const expected = "1011284"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func BenchmarkDay08_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay08_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
