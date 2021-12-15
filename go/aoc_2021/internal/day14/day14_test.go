package day14

import "testing"

func TestDay14_Part1(t *testing.T) {
	d := &day14{}
	d.Open()
	const expected = "2947"
	answer := d.Part1()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
	d.Close()
}

func TestDay14_Part2(t *testing.T) {
	d := &day14{}
	d.Open()
	const expected = "3232426226464"
	answer := d.Part2()
	if answer != expected {
		t.Fatalf("Expected %s but got %s", expected, answer)
	}
	d.Close()
}

func BenchmarkDay14_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay14_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
