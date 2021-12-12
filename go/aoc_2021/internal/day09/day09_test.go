package day09

import "testing"

func TestDay09_Part1(t *testing.T) {
	d := New()
	d.Open()
	const expected = "465"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}
func TestDay09_Part2(t *testing.T) {
	d := New()
	d.Open()
	const expected = "1269555"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
}

func BenchmarkDay09_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay09_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
