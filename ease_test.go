package geo

import "testing"

func TestLerp(t *testing.T) {
	cases := []struct {
		a, b, t float64
		want    float64
	}{
		{1, 5, 0, 1},
		{1, 5, 1, 5},
		{1, 5, 0.5, 3},
	}

	for i, c := range cases {
		got := Lerp(c.a, c.b, c.t)
		if got != c.want {
			t.Errorf("case %d: got %f, want %f", i, got, c.want)
		}
	}
}

func TestLerpVec(t *testing.T) {
	cases := []struct {
		a, b Vec
		t    float64
		want Vec
	}{
		{VecXY(1, 5), VecXY(5, 15), 0, VecXY(1, 5)},
		{VecXY(1, 5), VecXY(5, 15), 1, VecXY(5, 15)},
		{VecXY(1, 5), VecXY(5, 15), 0.5, VecXY(3, 10)},
	}

	for i, c := range cases {
		got := LerpVec(c.a, c.b, c.t)
		if got != c.want {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}
	}
}
