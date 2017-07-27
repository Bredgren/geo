package geo

import "testing"

func TestRectInflate(t *testing.T) {
	cases := []struct {
		dw, dh  float64
		r, want Rect
	}{
		{2, 2, Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 0, W: 7, H: 7}},
		{-1, -1, Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 1.5, Y: 1.5, W: 4, H: 4}},
	}

	for i, c := range cases {
		got := c.r.Inflated(c.dw, c.dh)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		c.r.Inflate(c.dw, c.dh)
		if c.r != c.want {
			t.Errorf("IP case %d: got %#v, want %#v", i, c.r, c.want)
		}
	}
}

func TestRectClamp(t *testing.T) {
	bounds := Rect{X: 1, Y: 1, W: 5, H: 5}
	cases := []struct {
		bounds, r, want Rect
	}{
		{bounds, Rect{X: 0, Y: 0, W: 1, H: 1}, Rect{X: 1, Y: 1, W: 1, H: 1}},
		{bounds, Rect{X: 7, Y: 6, W: 1, H: 1}, Rect{X: 5, Y: 5, W: 1, H: 1}},
		{bounds, Rect{X: 7, Y: 6, W: 7, H: 7}, Rect{X: 0, Y: 0, W: 7, H: 7}},
	}

	for i, c := range cases {
		got := c.r.Clamped(c.bounds)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		c.r.Clamp(c.bounds)
		if c.r != c.want {
			t.Errorf("IP case %d: got %#v, want %#v", i, c.r, c.want)
		}
	}
}

func TestRectIntersect(t *testing.T) {
	cases := []struct {
		r1, r2, want Rect
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 0, W: 2, H: 3}, Rect{X: 1, Y: 1, W: 1, H: 2}},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 2, Y: 3, W: 4, H: 5}, Rect{X: 2, Y: 3, W: 4, H: 3}},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 2, Y: 2, W: 2, H: 2}, Rect{X: 2, Y: 2, W: 2, H: 2}},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 6, Y: 6, W: 2, H: 2}, Rect{X: 6, Y: 6, W: 0, H: 0}},
	}

	for i, c := range cases {
		got := c.r1.Intersect(c.r2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		got = c.r2.Intersect(c.r1)
		if got != c.want {
			t.Errorf("reverse case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestRectUnion(t *testing.T) {
	cases := []struct {
		r1, r2, want Rect
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 0, W: 1, H: 1}, Rect{X: 0, Y: 0, W: 6, H: 6}},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 4, Y: 3, W: 3, H: 3}, Rect{X: 1, Y: 1, W: 6, H: 5}},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 2, Y: 2, W: 2, H: 2}, Rect{X: 1, Y: 1, W: 5, H: 5}},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 7, Y: 6, W: 7, H: 7}, Rect{X: 1, Y: 1, W: 13, H: 12}},
	}

	for i, c := range cases {
		got := c.r1.Unioned(c.r2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		got = c.r2.Unioned(c.r1)
		if got != c.want {
			t.Errorf("reverse case %d: got %#v, want %#v", i, got, c.want)
		}

		c.r1.Union(c.r2)
		if c.r1 != c.want {
			t.Errorf("IP case %d: got %#v, want %#v", i, c.r1, c.want)
		}
	}
}

func TestRectUnionAll(t *testing.T) {
	cases := []struct {
		rs   []Rect
		want Rect
	}{
		{
			rs: []Rect{
				Rect{X: 1, Y: 1, W: 5, H: 5},
				Rect{X: 0, Y: 2, W: 3, H: 6},
				Rect{X: 4, Y: -1, W: 4, H: 4},
			},
			want: Rect{X: 0, Y: -1, W: 8, H: 9},
		},
	}

	for i, c := range cases {
		got := RectUnion(c.rs)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestRectFit(t *testing.T) {
	rect1 := Rect{X: 1, Y: 1, W: 6, H: 3}
	rect2 := Rect{X: 2, Y: 2, W: 3, H: 6}
	rect3 := Rect{X: 3, Y: 3, W: 8, H: 2}
	rect4 := Rect{X: 4, Y: 4, W: 2, H: 8}

	cases := []struct {
		r1, r2, want Rect
	}{
		{rect1, rect2, Rect{X: 2, Y: 2, W: 3, H: 1.5}},
		{rect1, rect3, Rect{X: 3, Y: 3, W: 4, H: 2}},
		{rect1, rect4, Rect{X: 4, Y: 4, W: 2, H: 1}},

		{rect2, rect1, Rect{X: 2, Y: 1, W: 1.5, H: 3}},
		{rect2, rect3, Rect{X: 3, Y: 3, W: 1, H: 2}},
		{rect2, rect4, Rect{X: 4, Y: 4, W: 2, H: 4}},

		{rect3, rect1, Rect{X: 1, Y: 4 - 6.0/4.0, W: 6, H: 6.0 / 4.0}},
		{rect3, rect2, Rect{X: 2, Y: 3, W: 3, H: 3.0 / 4.0}},
		{rect3, rect4, Rect{X: 4, Y: 4, W: 2, H: 2.0 / 4.0}},

		{rect4, rect1, Rect{X: 4, Y: 1, W: 3.0 / 4.0, H: 3}},
		{rect4, rect2, Rect{X: 5 - 6.0/4.0, Y: 2, W: 6.0 / 4.0, H: 6}},
		{rect4, rect3, Rect{X: 4, Y: 3, W: 2.0 / 4.0, H: 2.0}},
	}

	for i, c := range cases {
		got := c.r1.Fitted(c.r2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		c.r1.Fit(c.r2)
		if c.r1 != c.want {
			t.Errorf("IP case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestRectNormalize(t *testing.T) {
	cases := []struct {
		r, want Rect
	}{
		{Rect{X: 1, Y: 1, W: -4, H: -2}, Rect{X: -3, Y: -1, W: 4, H: 2}},
	}

	for i, c := range cases {
		c.r.Normalize()
		if c.r != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, c.r, c.want)
		}
	}
}

func TestRectContains(t *testing.T) {
	cases := []struct {
		r1, r2 Rect
		want   bool
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 2, Y: 2, W: 5, H: 2}, false},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 2, Y: 2, W: 4, H: 2}, true},
	}

	for i, c := range cases {
		got := c.r1.Contains(c.r2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestRectCollidePoint(t *testing.T) {
	cases := []struct {
		r    Rect
		x, y float64
		want bool
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5}, 0, 0, false},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, 1, 1, true},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, 4, 4, true},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, 5, 5, true},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, 6, 6, false},
	}

	for i, c := range cases {
		got := c.r.CollidePoint(c.x, c.y)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestRectCollideRect(t *testing.T) {
	cases := []struct {
		r1, r2 Rect
		want   bool
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 0, W: 7, H: 1}, false},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 0, W: 1, H: 7}, false},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 6, Y: 0, W: 2, H: 7}, false},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 6, W: 7, H: 2}, false},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 0, Y: 0, W: 2, H: 2}, true},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, Rect{X: 5, Y: 5, W: 2, H: 2}, true},
	}

	for i, c := range cases {
		got := c.r1.CollideRect(c.r2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}
