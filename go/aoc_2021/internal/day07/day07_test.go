package day07

import "testing"

func TestDay07_Part1(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part1()
	if answer != "356958" {
		t.Fatalf("Expected 356958 but got %s", answer)
	}
}

func TestDay07_Part2(t *testing.T) {
	d := New()
	d.Open()
	answer := d.Part2()
	if answer != "105461913" {
		t.Fatalf("Expected 105461913 but got %s", answer)
	}
}

func BenchmarkDay07_Part1_Bruteforce(b *testing.B) {
	d := &day07{bin: false}
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay07_Part1_Binary(b *testing.B) {
	d := &day07{bin: true}
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay07_Part2_Bruteforce(b *testing.B) {
	d := &day07{bin: false}
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}

func BenchmarkDay07_Part2_Binary(b *testing.B) {
	d := &day07{bin: true}
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
