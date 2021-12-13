package day13

import (
	"strings"
	"testing"
)

func TestDay13_Part1(t *testing.T) {
	d := New()
	d.Open()
	const expected = "682"
	actual := d.Part1()
	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
func TestDay13_Part2(t *testing.T) {
	d := New()
	d.Open()
	expected := `
@@@@__@@___@@__@__@_@@@__@@@@_@__@_@@@@
@____@__@_@__@_@__@_@__@____@_@__@_@___
@@@__@__@_@____@__@_@__@___@__@@@@_@@@_
@____@@@@_@_@@_@__@_@@@___@___@__@_@___
@____@__@_@__@_@__@_@_@__@____@__@_@___
@____@__@__@@@__@@__@__@_@@@@_@__@_@@@@
`
	actual := d.Part2()
	if strings.Trim(actual, " \n") != strings.Trim(expected, " \n") {
		t.Fatalf("Expected '%s' but got '%s'", expected, actual)
	}
}

func BenchmarkDay13_Part1(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}

func BenchmarkDay13_Part2(b *testing.B) {
	d := New()
	d.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Part2()
	}
}
