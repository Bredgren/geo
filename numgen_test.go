package geo

import (
	"math"
	"math/rand"
	"testing"
)

func isBetween(n, a, b float64) bool {
	if a > b {
		a, b = b, a
	}
	return a <= n && n < b || n == a && n == b
}

func isBetweenErr(n, a, b, e float64) bool {
	if a > b {
		a, b = b, a
	}
	return a-e <= n && n < b+e
}

func TestConsNum(t *testing.T) {
	trials := 1000
	n := rand.Float64()
	gen := ConstNum(n)
	for i := 0; i < trials; i++ {
		got := gen()
		if got != n {
			t.Errorf("got %f, want %f", got, n)
		}
	}
}

func TestRandNum(t *testing.T) {
	trials := 1000
	cases := []struct {
		a, b float64
	}{
		{0, 0},
		{0, 1},
		{2, 2},
		{3, 2},
		{-1, 1},
		{1, -1},
		{-1, -3},
		{-3, -1},
	}
	for i, c := range cases {
		gen := RandNum(c.a, c.b)
		for j := 0; j < trials; j++ {
			got := gen()
			if !isBetween(got, c.a, c.b) {
				t.Errorf("case %d: trial %d: got %f, want between %f and %f", i, j, got, c.a, c.b)
			}
		}
	}
}

func TestRandRadius(t *testing.T) {
	trials := 1000
	cases := []struct {
		a, b float64
	}{
		{0, 0},
		{0, 1},
		{2, 2},
		{3, 2},
	}
	for i, c := range cases {
		gen := RandRadius(c.a, c.b)
		for j := 0; j < trials; j++ {
			got := gen()
			if !isBetween(got, c.a, c.b) {
				t.Errorf("case %d: trial %d: got %f, want between %f and %f", i, j, got, c.a, c.b)
			}
		}
	}
}

func TestStaticVec(t *testing.T) {
	trials := 10
	v := RandVec()
	gen := StaticVec(v)
	for i := 0; i < trials; i++ {
		got := gen()
		if !got.Equals(v, e) {
			t.Errorf("got %s, want %s", got, v)
		}
	}
}

func TestDynamicVec(t *testing.T) {
	v := Vec{}
	dyn := DynamicVec(&v)
	v2 := dyn()
	if !v.Equals(v2, e) {
		t.Error(v, v2)
	}
	v.X = 10
	v3 := dyn()
	if !v.Equals(v3, e) {
		t.Error(v, v3)
	}
}

func TestOffestVec(t *testing.T) {
	trials := 1000
	v := RandVec()
	for i := 0; i < trials; i++ {
		offset := RandVec()
		gen := OffsetVec(StaticVec(v), StaticVec(offset))
		got := gen()
		if !got.Equals(v.Plus(offset), e) {
			t.Errorf("got %s, want %s", got, v)
		}
	}
}

func TestRandVecCircle(t *testing.T) {
	trials := 1000
	cases := []struct {
		a, b float64
	}{
		{0, 0},
		{0, 1},
		{1, 1},
		{0, 5},
		{5, 3},
		{5, 5},
	}

	for i, c := range cases {
		vecGen := RandVecCircle(c.a, c.b)
		for l := 0; l < trials; l++ {
			v := vecGen()
			if !isBetweenErr(v.Len(), c.a, c.b, e) {
				t.Errorf("case %d: trial %d: got %f, want between %f and %f, v: %s",
					i, l, v.Len(), c.a, c.b, v)
			}
		}
	}
}

func TestRandVecArc(t *testing.T) {
	trials := 1000
	cases := []struct {
		r1, r2, rad1, rad2 float64
	}{
		{0, 0, -math.Pi / 4, math.Pi / 4},
		{0, 1, -math.Pi / 4, math.Pi / 4},
		{1, 1, -math.Pi / 4, math.Pi / 4},
		{0, 5, math.Pi / 4, -math.Pi / 4},
		{5, 3, math.Pi / 4, -math.Pi / 4},
		{5, 5, math.Pi / 4, -math.Pi / 4},
	}

	for i, c := range cases {
		vecGen := RandVecArc(c.r1, c.r2, c.rad1, c.rad2)
		for l := 0; l < trials; l++ {
			v := vecGen()
			len := v.Len()
			rad := v.Angle()
			if !isBetweenErr(len, c.r1, c.r2, e) || !isBetween(rad, c.rad1, c.rad2) {
				t.Errorf("case %d: trial %d: got len %f, rad %f, want len between %f and %f rad between %f and %f",
					i, l, len, rad, c.r1, c.r2, c.rad1, c.rad2)
			}
		}
	}
}

func TestRandVecRect(t *testing.T) {
	trials := 1000
	rect := Rect{
		X: rand.Float64()*100 - 50,
		Y: rand.Float64()*100 - 50,
		W: rand.Float64() * 100,
		H: rand.Float64() * 100,
	}

	vecGen := RandVecRect(rect)
	for l := 0; l < trials; l++ {
		v := vecGen()
		if !rect.CollidePoint(v.X, v.Y) {
			t.Errorf("trial %d: %s, %s", l, rect, v)
		}
	}
}

func TestRandVecRects(t *testing.T) {
	zeroGen := RandVecRects([]Rect{})
	v := zeroGen()
	if !v.Equals(Vec{}, e) {
		t.Errorf("no rects: got %s, want %s", v, Vec{})
	}

	trials := 1000
	numRects := rand.Intn(19) + 1

	rects := []Rect{}
	for i := 0; i < numRects; i++ {
		rects = append(rects, Rect{
			X: rand.Float64()*100 - 50,
			Y: rand.Float64()*100 - 50,
			W: rand.Float64() * 100,
			H: rand.Float64() * 100,
		})
	}
	t.Logf("rects: %v", rects)

	vecGen := RandVecRects(rects)
	for l := 0; l < trials; l++ {
		v := vecGen()
		collides := false
		for _, r := range rects {
			if r.CollidePoint(v.X, v.Y) {
				collides = true
				break
			}
		}
		if !collides {
			t.Errorf("trial %d: %s", l, v)
		}
	}
}
