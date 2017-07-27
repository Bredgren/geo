package geo

import "testing"

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
