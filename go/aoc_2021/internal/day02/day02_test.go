package day02

import "testing"

func BenchmarkDay02_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := New()
		d.Open()
		d.Part1()
	}
}

func BenchmarkDay02_Part2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		d := New()
		d.Open()
		d.Part2()
	}
}
