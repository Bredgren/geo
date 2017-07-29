package geo

import "testing"

func TestMakeRect(t *testing.T) {
	want := Rect{X: 1, Y: 2, W: 2, H: 3}
	got := RectXYWH(1, 2, 2, 3)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = RectCorners(1, 2, 3, 5)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = RectCornersVec(Vec{X: 1, Y: 2}, Vec{X: 3, Y: 5})
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	r := RectCorners(-1.2, -3.4, 5.6, 7.8)
	ir := r.Rectangle()
	got = Rectangle(ir)
	if got.X != -1 || got.Y != -3 || got.Right() != 5 || got.Bottom() != 7 {
		t.Errorf("got %s r: %g, b: %g, want %s", got, got.Right(), got.Bottom(), ir)
	}

	got = RectWH(2, 3)
	got.SetTopLeft(1, 2)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = RectVWH(VecXY(1, 2), 2, 3)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = RectVSize(VecXY(1, 2), VecXY(2, 3))
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRectString(t *testing.T) {
	r := Rect{X: -1.2, Y: 3.4, W: 5.6, H: 7.8}
	got := r.String()
	want := "Rect(-1.2, 3.4, w5.6, h7.8)"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRectGetSetFuncs(t *testing.T) {
	r := Rect{X: -4, Y: 3, W: 10, H: 6}

	check := func(name string, fn func() float64, want float64) {
		got := fn()
		if got != want {
			t.Errorf("%s: got %f, want %f", name, got, want)
		}
	}

	check2 := func(name string, fn func() (float64, float64), want1, want2 float64) {
		got1, got2 := fn()
		if got1 != want1 || got2 != want2 {
			t.Errorf("%s: got %f, %f, want %f, %f", name, got1, got2, want1, want2)
		}
	}

	check("top", r.Top, 3)
	check("left", r.Left, -4)
	check("right", r.Right, 6)
	check("bottom", r.Bottom, 9)

	check2("size", r.Size, 10, 6)

	check2("top left", r.TopLeft, -4, 3)
	check2("top mid", r.TopMid, 1, 3)
	check2("top right", r.TopRight, 6, 3)
	check2("left mid", r.LeftMid, -4, 6)
	check2("mid", r.Mid, 1, 6)
	check("mid x", r.MidX, 1)
	check("mid y", r.MidY, 6)
	check2("right mid", r.RightMid, 6, 6)
	check2("bottom left", r.BottomLeft, -4, 9)
	check2("bottom mid", r.BottomMid, 1, 9)
	check2("bottom right", r.BottomRight, 6, 9)

	check("area", r.Area, 60)

	r.SetTop(1)
	check("set top", r.Top, 1)
	r.SetLeft(-2)
	check("set left", r.Left, -2)
	r.SetRight(10)
	check("set right", r.Right, 10)
	r.SetBottom(11)
	check("set bottom", r.Bottom, 11)

	r.SetSize(6, 10)
	check2("set size", r.Size, 6, 10)

	r.SetTopLeft(2, -3)
	check2("set top left", r.TopLeft, 2, -3)
	r.SetTopMid(3, -4)
	check2("set top mid", r.TopMid, 3, -4)
	r.SetTopRight(7, -4)
	check2("set top right", r.TopRight, 7, -4)
	r.SetLeftMid(-6, 0)
	check2("set left mid", r.LeftMid, -6, 0)
	r.SetMid(-1, -1)
	check2("set mid", r.Mid, -1, -1)
	r.SetMidX(2)
	check("set mid x", r.MidX, 2)
	r.SetMidY(4)
	check("set mid y", r.MidY, 4)
	r.SetRightMid(-1, 5)
	check2("set right mid", r.RightMid, -1, 5)
	r.SetBottomLeft(5, -5)
	check2("set bottom left", r.BottomLeft, 5, -5)
	r.SetBottomMid(-4, 4)
	check2("set bottom mid", r.BottomMid, -4, 4)
	r.SetBottomRight(-6, -10)
	check2("set bottom right", r.BottomRight, -6, -10)

	r.Move(10, -12)
	check2("move", r.BottomRight, 4, -22)
	got := r.Moved(-5, 6)
	check2("moved", got.BottomRight, -1, -16)
}

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
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		c.r.Inflate(c.dw, c.dh)
		if c.r != c.want {
			t.Errorf("IP case %d: got %s, want %s", i, c.r, c.want)
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
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		c.r.Clamp(c.bounds)
		if c.r != c.want {
			t.Errorf("IP case %d: got %s, want %s", i, c.r, c.want)
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
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		got = c.r2.Intersect(c.r1)
		if got != c.want {
			t.Errorf("reverse case %d: got %s, want %s", i, got, c.want)
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
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		got = c.r2.Unioned(c.r1)
		if got != c.want {
			t.Errorf("reverse case %d: got %s, want %s", i, got, c.want)
		}

		c.r1.Union(c.r2)
		if c.r1 != c.want {
			t.Errorf("IP case %d: got %s, want %s", i, c.r1, c.want)
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
		{
			rs:   []Rect{},
			want: Rect{},
		},
	}

	for i, c := range cases {
		got := RectUnion(c.rs)
		if got != c.want {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
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
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		c.r1.Fit(c.r2)
		if c.r1 != c.want {
			t.Errorf("IP case %d: got %s, want %s", i, got, c.want)
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
		got := c.r.Normalized()
		if got != c.want {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}
		c.r.Normalize()
		if c.r != c.want {
			t.Errorf("case %d: got %s, want %s", i, c.r, c.want)
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
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
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
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
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
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}

func TestRectCollideList(t *testing.T) {
	cases := []struct {
		r    Rect
		rs   []Rect
		want int
		ok   bool
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5},
			[]Rect{
				Rect{X: 0, Y: 0, W: 7, H: 1},
				Rect{X: 0, Y: 0, W: 1, H: 7},
				Rect{X: 6, Y: 0, W: 2, H: 7},
				Rect{X: 0, Y: 6, W: 7, H: 2},
				Rect{X: 0, Y: 0, W: 2, H: 2},
				Rect{X: 5, Y: 5, W: 2, H: 2},
			},
			5, true,
		},
		{Rect{X: 1, Y: 1, W: 5, H: 5},
			[]Rect{
				Rect{X: 0, Y: 0, W: 7, H: 1},
				Rect{X: 0, Y: 0, W: 1, H: 7},
				Rect{X: 6, Y: 0, W: 2, H: 7},
				Rect{X: 0, Y: 6, W: 7, H: 2},
			},
			0, false,
		},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, []Rect{}, 0, false},
	}

	for i, c := range cases {
		got, ok := c.r.CollideRectList(c.rs)
		if c.ok == ok && got != c.want {
			t.Errorf("case %d: got %d, want %d", i, got, c.want)
		}
	}
}

func TestRectCollideListAll(t *testing.T) {
	cases := []struct {
		r    Rect
		rs   []Rect
		want []int
	}{
		{Rect{X: 1, Y: 1, W: 5, H: 5},
			[]Rect{
				Rect{X: 0, Y: 0, W: 7, H: 1},
				Rect{X: 0, Y: 0, W: 1, H: 7},
				Rect{X: 6, Y: 0, W: 2, H: 7},
				Rect{X: 0, Y: 0, W: 2, H: 2},
				Rect{X: 0, Y: 6, W: 7, H: 2},
				Rect{X: 5, Y: 5, W: 2, H: 2},
			},
			[]int{3, 5},
		},
		{Rect{X: 1, Y: 1, W: 5, H: 5},
			[]Rect{
				Rect{X: 0, Y: 0, W: 7, H: 1},
				Rect{X: 0, Y: 0, W: 1, H: 7},
				Rect{X: 6, Y: 0, W: 2, H: 7},
				Rect{X: 0, Y: 6, W: 7, H: 2},
			},
			[]int{},
		},
		{Rect{X: 1, Y: 1, W: 5, H: 5}, []Rect{}, []int{}},
	}

	listEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, n := range a {
			if n != b[i] {
				return false
			}
		}
		return true
	}

	for i, c := range cases {
		got := c.r.CollideRectListAll(c.rs)
		if !listEqual(got, c.want) {
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}
