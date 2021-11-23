package day05

import "testing"

func TestParser(t *testing.T) {
	cases := []struct {
		line    string
		r, c, s int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	d := &day05{}
	for _, tc := range cases {
		t.Run(tc.line, func(t *testing.T) {
			d.loadPasses([]string{tc.line})
			if d.seats[0].r != tc.r {
				t.Fatalf("expected row to be %d but got %d", tc.r, d.seats[0].r)
			}
			if d.seats[0].c != tc.c {
				t.Fatalf("expected column to be %d but got %d", tc.c, d.seats[0].c)
			}
			if d.seats[0].id != tc.s {
				t.Fatalf("expected id to be %d but got %d", tc.s, d.seats[0].id)
			}
		})
	}
}
