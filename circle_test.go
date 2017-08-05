package geo

import (
	"math"
	"testing"
)

func TestMakeCircle(t *testing.T) {
	want := Circle{X: 1, Y: 2, R: 3}
	got := CircleXYR(1, 2, 3)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = CircleXY(1, 2)
	got.R = 3
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = CircleVecR(VecXY(1, 2), 3)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = CircleVec(VecXY(1, 2))
	got.R = 3
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	got = CircleR(3)
	got.X = 1
	got.Y = 2
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	r := want.InscribedRect()
	got = CircleCircumscribe(r)
	if !got.Equals(want, e) {
		t.Errorf("got %s, want %s", got, want)
	}

	r = want.BoundingRect()
	got = CircleInscribe(r)
	if !got.Equals(want, e) {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCircleString(t *testing.T) {
	c := CircleXYR(-1.2, 3.4, 5.6)
	got := c.String()
	want := "Circle(-1.2, 3.4, r5.6)"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCircleGetSetFuncs(t *testing.T) {
	c := CircleXYR(-4, 3, 10)

	check(t, "top", c.Top, -7)
	check(t, "left", c.Left, -14)
	check(t, "right", c.Right, 6)
	check(t, "bottom", c.Bottom, 13)

	check2(t, "top left", c.TopLeft, -14, -7)
	check2(t, "top mid", c.TopMid, -4, -7)
	check2(t, "top right", c.TopRight, 6, -7)
	check2(t, "left mid", c.LeftMid, -14, 3)
	check2(t, "mid", c.Mid, -4, 3)
	checkVec(t, "pos", c.Pos, VecXY(-4, 3))
	check2(t, "right mid", c.RightMid, 6, 3)
	check2(t, "bottom left", c.BottomLeft, -14, 13)
	check2(t, "bottom mid", c.BottomMid, -4, 13)
	check2(t, "bottom right", c.BottomRight, 6, 13)

	check(t, "circumference", c.Circumference, 2*math.Pi*10)
	check(t, "diameter", c.Diameter, 20)

	check3(t, "xyr", c.XYR, -4, 3, 10)
	gotV, gotR := c.PosR()
	if gotV != VecXY(-4, 3) || gotR != 10 {
		t.Errorf("posr: got %s, %f, want %s, %f", gotV, gotR, c.Pos(), c.R)
	}

	check(t, "area", c.Area, math.Pi*10*10)

	c.SetTop(1)
	check(t, "set top", c.Top, 1)
	c.SetLeft(-2)
	check(t, "set left", c.Left, -2)
	c.SetRight(10)
	check(t, "set right", c.Right, 10)
	c.SetBottom(11)
	check(t, "set bottom", c.Bottom, 11)

	c.SetTopLeft(2, -3)
	check2(t, "set top left", c.TopLeft, 2, -3)
	c.SetTopMid(3, -4)
	check2(t, "set top mid", c.TopMid, 3, -4)
	c.SetTopRight(7, -4)
	check2(t, "set top right", c.TopRight, 7, -4)
	c.SetLeftMid(-6, 0)
	check2(t, "set left mid", c.LeftMid, -6, 0)
	c.SetMid(-1, -1)
	check2(t, "set mid", c.Mid, -1, -1)
	c.SetPos(VecXY(-10, 20))
	checkVec(t, "set pos", c.Pos, VecXY(-10, 20))
	c.SetRightMid(-1, 5)
	check2(t, "set right mid", c.RightMid, -1, 5)
	c.SetBottomLeft(5, -5)
	check2(t, "set bottom left", c.BottomLeft, 5, -5)
	c.SetBottomMid(-4, 4)
	check2(t, "set bottom mid", c.BottomMid, -4, 4)
	c.SetBottomRight(-6, -10)
	check2(t, "set bottom right", c.BottomRight, -6, -10)

	c.Move(10, -12)
	check2(t, "move", c.BottomRight, 4, -22)
	got := c.Moved(-5, 6)
	check2(t, "moved", got.BottomRight, -1, -16)

	c.SetCircumference(12)
	check(t, "circumference", c.Circumference, 12)

	c.SetDiameter(10)
	check(t, "diameter", c.Diameter, 10)
}

func TestCircleInflate(t *testing.T) {
	cases := []struct {
		dr      float64
		c, want Circle
	}{
		{2, CircleXYR(1, 2, 3), CircleXYR(1, 2, 5)},
		{-1, CircleXYR(1, 2, 3), CircleXYR(1, 2, 2)},
	}

	for i, c := range cases {
		got := c.c.Inflated(c.dr)
		if got != c.want {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		c.c.Inflate(c.dr)
		if c.c != c.want {
			t.Errorf("IP case %d: got %s, want %s", i, c.c, c.want)
		}
	}
}

func TestCircleNormalize(t *testing.T) {
	cases := []struct {
		c, want Circle
	}{
		{CircleXYR(1, 2, 3), CircleXYR(1, 2, 3)},
		{CircleXYR(1, 2, -3), CircleXYR(1, 2, 3)},
	}

	for i, c := range cases {
		got := c.c.Normalized()
		if got != c.want {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		c.c.Normalize()
		if c.c != c.want {
			t.Errorf("IP case %d: got %s, want %s", i, c.c, c.want)
		}
	}
}

func TestCircleContains(t *testing.T) {
	cases := []struct {
		c1, c2 Circle
		want   bool
	}{
		{CircleXYR(0, 1, 2), CircleXYR(0, 5, 3), false},
		{CircleXYR(0, 1, 2), CircleXYR(0, 2, 5), false},
		{CircleXYR(0, 2, 5), CircleXYR(0, 1, 2), true},
	}

	for i, c := range cases {
		got := c.c1.Contains(c.c2)
		if got != c.want {
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}

func TestCircleCollideCircle(t *testing.T) {
	cases := []struct {
		c1, c2 Circle
		want   bool
	}{
		{CircleXYR(0, 1, 2), CircleXYR(0, 5, 3), true},
		{CircleXYR(0, 1, 2), CircleXYR(0, 2, 5), true},
		{CircleXYR(0, 2, 5), CircleXYR(0, 1, 2), true},
		{CircleXYR(0, 1, 2), CircleXYR(0, 4, 1.1), true},
		{CircleXYR(0, 1, 2), CircleXYR(0, 4, 1), false},
		{CircleXYR(1, 1, 5), CircleXYR(0, 0, 7), true},
	}

	for i, c := range cases {
		got := c.c1.CollideCircle(c.c2)
		if got != c.want {
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}

func TestCircleCollidePoint(t *testing.T) {
	cases := []struct {
		c    Circle
		p    Vec
		want bool
	}{
		{CircleXYR(0, 1, 2), Vec0, true},
		{CircleXYR(0, 1, 2), VecXY(0, -1), false},
		{CircleXYR(0, 1, 2), VecXY(0, 2), true},
		{CircleXYR(0, 1, 2), VecXY(0, 3), false},
		{CircleXYR(0, 1, 2), VecXY(0, 2.99), true},
	}

	for i, c := range cases {
		got := c.c.CollidePoint(c.p.XY())
		if got != c.want {
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}

func TestCircleCollideList(t *testing.T) {
	cases := []struct {
		c    Circle
		cs   []Circle
		want int
		ok   bool
	}{
		{CircleXYR(1, 1, 5),
			[]Circle{
				CircleXYR(10, 10, 7),
				CircleXYR(-7, 0, 1),
				CircleXYR(7, 0, 1),
				CircleXYR(0, 8, 2),
				CircleXYR(2, 2, 2),
				CircleXYR(-5, -5, 2),
				CircleXYR(0, 0, 2),
			},
			4, true,
		},
		{CircleXYR(1, 1, 5),
			[]Circle{
				CircleXYR(10, 10, 7),
				CircleXYR(-7, 0, 1),
				CircleXYR(7, 0, 1),
				CircleXYR(0, 8, 2),
				CircleXYR(-5, -5, 2),
			},
			0, false,
		},
		{CircleXYR(1, 1, 5), []Circle{}, 0, false},
	}

	for i, c := range cases {
		got, ok := c.c.CollideCircleList(c.cs)
		if c.ok == ok && got != c.want {
			t.Errorf("case %d: got %d, want %d", i, got, c.want)
		}
	}
}

func TestCircleCollideListAll(t *testing.T) {
	cases := []struct {
		c    Circle
		cs   []Circle
		want []int
	}{
		{CircleXYR(1, 1, 5),
			[]Circle{
				CircleXYR(10, 10, 7),
				CircleXYR(-7, 0, 1),
				CircleXYR(7, 0, 1),
				CircleXYR(0, 8, 2),
				CircleXYR(2, 2, 2),
				CircleXYR(-5, -5, 2),
				CircleXYR(0, 0, 2),
			},
			[]int{4, 6},
		},
		{CircleXYR(1, 1, 5),
			[]Circle{
				CircleXYR(10, 10, 7),
				CircleXYR(-7, 0, 1),
				CircleXYR(7, 0, 1),
				CircleXYR(0, 8, 2),
				CircleXYR(-5, -5, 2),
			},
			[]int{},
		},
		{CircleXYR(1, 1, 5), []Circle{}, []int{}},
	}

	for i, c := range cases {
		got := c.c.CollideCircleListAll(c.cs)
		if !intListEqual(got, c.want) {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestCircleUnion(t *testing.T) {
	cases := []struct {
		c1, c2, want Circle
	}{
		{CircleXYR(1, 1, 5), CircleXYR(1, 1, 5), CircleXYR(1, 1, 5)},
		{CircleXYR(-2, 1, 5), CircleXYR(2, 1, 5), CircleXYR(0, 1, 7)},
		{CircleXYR(-2, 1, 5), CircleXYR(3, 1, 5), CircleXYR(0.5, 1, 7.5)},
	}

	for i, c := range cases {
		got := c.c1.Unioned(c.c2)
		if got != c.want {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}

		c.c1.Union(c.c2)
		if c.c1 != c.want {
			t.Errorf("case %d: got %s, want %s", i, c.c1, c.want)
		}
	}
}

func TestPointAt(t *testing.T) {
	cases := []struct {
		c    Circle
		rad  float64
		want Vec
	}{
		{CircleXYR(1, 0, 5), 0, VecXY(6, 0)},
		{CircleXYR(1, 0, 5), math.Pi / 2, VecXY(1, -5)},
		{CircleXYR(1, 0, 5), math.Pi, VecXY(-4, 0)},
		{CircleXYR(1, -1, 5), math.Pi * 3 / 2, VecXY(1, 4)},
	}

	for i, c := range cases {
		got := VecXY(c.c.PointAt(c.rad))
		if !got.Equals(c.want, e) {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}
	}
}

func TestCircleCollideRect(t *testing.T) {
	cases := []struct {
		c    Circle
		r    Rect
		want bool
	}{
		{CircleXYR(-2, 0, 1), RectXYWH(0, 0, 7, 1), false},
		{CircleXYR(-1, 0, 1), RectXYWH(0, 0, 7, 1), false},
		{CircleXYR(-0.9, 0, 1), RectXYWH(0, 0, 7, 1), true},
		{CircleXYR(5, 0, 1), RectXYWH(0, 0, 7, 1), true},
		{CircleXYR(5, -2, 1), RectXYWH(0, 0, 7, 1), false},
		{CircleXYR(5, -2, 2), RectXYWH(0, 0, 7, 1), false},
		{CircleXYR(5, -2, 2.1), RectXYWH(0, 0, 7, 1), true},
		{CircleXYR(8, 3, 1), RectXYWH(0, 0, 7, 1), false},
		{CircleXYR(8, 3, 2), RectXYWH(0, 0, 7, 1), false},
		{CircleXYR(8, 3, 3), RectXYWH(0, 0, 7, 1), true},
	}

	for i, c := range cases {
		got := c.c.CollideRect(c.r)
		if got != c.want {
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}

func TestCircleCollideRectList(t *testing.T) {
	cases := []struct {
		c    Circle
		rs   []Rect
		want int
		ok   bool
	}{
		{CircleXYR(4, 3, 2),
			[]Rect{
				RectXYWH(2, -3, 4, 2),
				RectXYWH(6, 3, 3, 3),
				RectXYWH(-1, 3, 4, 3),
				RectXYWH(-5, -3, 2, 4),
				RectXYWH(4, 0, 3, 2),
			},
			2, true,
		},
		{CircleXYR(4, 3, 2),
			[]Rect{
				RectXYWH(2, -3, 4, 2),
				RectXYWH(6, 3, 3, 3),
				RectXYWH(-5, -3, 2, 4),
			},
			0, false,
		},
		{CircleXY(1, 1), []Rect{}, 0, false},
	}

	for i, c := range cases {
		got, ok := c.c.CollideRectList(c.rs)
		if c.ok == ok && got != c.want {
			t.Errorf("case %d: got %d, want %d", i, got, c.want)
		}
	}
}

func TestCircleCollideRectListAll(t *testing.T) {
	cases := []struct {
		c    Circle
		rs   []Rect
		want []int
	}{
		{CircleXYR(4, 3, 2),
			[]Rect{
				RectXYWH(2, -3, 4, 2),
				RectXYWH(6, 3, 3, 3),
				RectXYWH(-1, 3, 4, 3),
				RectXYWH(-5, -3, 2, 4),
				RectXYWH(4, 0, 3, 2),
			},
			[]int{2, 4},
		},
		{CircleXYR(4, 3, 2),
			[]Rect{
				RectXYWH(2, -3, 4, 2),
				RectXYWH(6, 3, 3, 3),
				RectXYWH(-5, -3, 2, 4),
			},
			[]int{},
		},
		{CircleXY(1, 1), []Rect{}, []int{}},
	}

	for i, c := range cases {
		got := c.c.CollideRectListAll(c.rs)
		if !intListEqual(got, c.want) {
			t.Errorf("case %d: got %v, want %v", i, got, c.want)
		}
	}
}
