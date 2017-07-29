package geo

import (
	"math"
	"testing"
)

func TestClamp(t *testing.T) {
	cases := []struct {
		n, a, b, want float64
	}{
		{0, 0, 0, 0},
		{0, 1, 2, 1},
		{0, 2, 1, 1},
		{3, 1, 2, 2},
		{3, 2, 1, 2},
		{1.2, 1, 2, 1.2},
		{1.2, 2, 1, 1.2},
	}
	for i, c := range cases {
		got := Clamp(c.n, c.a, c.b)
		if got != c.want {
			t.Errorf("case %d: got: %f, want: %f", i, got, c.want)
		}
	}
}

func TestMap(t *testing.T) {
	cases := []struct {
		n, a1, b1, a2, b2, want float64
	}{
		{0, 0, 1, 5, 10, 5},
		{1, 0, 1, 5, 10, 10},
		{0.5, 0, 1, 5, 11, 8},
		{1, 0, 0, 5, 11, math.Inf(1)},
		{1, 0, 2, 5, 5, 5},
	}
	for i, c := range cases {
		got := Map(c.n, c.a1, c.b1, c.a2, c.b2)
		if got != c.want {
			t.Errorf("case %d: got: %f, want: %f", i, got, c.want)
		}
	}
}
