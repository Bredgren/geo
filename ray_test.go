package geo

import (
	"math"
	"testing"
)

func TestRayString(t *testing.T) {
	r := Ray{VecXY(1.2, 3.4), VecXY(4.6, 7.8)}
	got := r.String()
	want := "Ray(Vec(1.2, 3.4), Vec(4.6, 7.8))"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRayAt(t *testing.T) {
	cases := []struct {
		r    Ray
		t    float64
		want Vec
	}{
		{Ray{VecXY(0, 0), VecXY(0, 1)}, 5, VecXY(0, 5)},
		{Ray{VecXY(0, 0), VecXY(-2, 0)}, 5, VecXY(-5, 0)},
		{Ray{VecXY(10, -5), VecXY(2, 2)}, 2 * math.Sqrt2, VecXY(12, -3)},
		{Ray{VecXY(10, -5), VecXY(0, 0)}, 5, VecXY(10, -5)},
	}

	for i, c := range cases {
		got := c.r.At(c.t)
		if !got.Equals(c.want, e) {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}
	}
}

func TestRayAngle(t *testing.T) {
	cases := []struct {
		origin Vec
		rad    float64
		want   Ray
	}{
		{VecXY(0, 0), 0, Ray{VecXY(0, 0), VecXY(1, 0)}},
		{VecXY(0, 0), math.Pi / 2, Ray{VecXY(0, 0), VecXY(0, -1)}},
		{VecXY(0, 0), math.Pi, Ray{VecXY(0, 0), VecXY(-1, 0)}},
		{VecXY(1, -1), math.Pi / 2, Ray{VecXY(1, -1), VecXY(0, -1)}},
	}

	for i, c := range cases {
		got := RayAngle(c.origin, c.rad)
		if !got.Origin.Equals(c.want.Origin, e) || !got.Direction.Equals(c.want.Direction, e) {
			t.Errorf("case %d: got %s, want %s", i, got, c.want)
		}
	}
}

func TestRayIntersectCircle(t *testing.T) {
	circ := CircleXYR(2, 2, 3)
	caseAngleT := VecXY(circ.PointAt(3 * math.Pi / 4)).Dist(VecXY(-4, -4))
	type res struct {
		tMin, tMax float64
		hit        bool
	}
	cases := []struct {
		r    Ray
		c    Circle
		want res
	}{
		// From the left
		{Ray{VecXY(-6, 2), VecXY(1, 0)}, circ, res{5, 11, true}},
		{Ray{VecXY(-6, 2), VecXY(-1, 0)}, circ, res{-11, -5, true}},
		{Ray{VecXY(-6, 2), VecXY(0, 1)}, circ, res{0, 0, false}},
		// From the right
		{Ray{VecXY(11, 2), VecXY(-1, 0)}, circ, res{6, 12, true}},
		{Ray{VecXY(11, 2), VecXY(1, 0)}, circ, res{-12, -6, true}},
		{Ray{VecXY(11, 2), VecXY(0, -1)}, circ, res{0, 0, false}},
		// Tangent
		{Ray{VecXY(5, 8), VecXY(0, -1)}, circ, res{0, 0, false}},
		// Angle
		{Ray{VecXY(-4, -4), VecXY(1, 1)}, circ, res{caseAngleT, caseAngleT + circ.Diameter(), true}},
	}

	for i, c := range cases {
		gotTMin, gotTMax, gotHit := c.r.IntersectCircle(c.c)
		if !fEqual(gotTMin, c.want.tMin) || !fEqual(gotTMax, c.want.tMax) || gotHit != c.want.hit {
			t.Errorf("case %d: gotTMin %f, gotTMax %f, gotHit %v, want %v", i, gotTMin, gotTMax,
				gotHit, c.want)
		}
	}
}

func TestRayIntersectRect(t *testing.T) {
	rect := RectXYWH(-1, -2, 6, 7)
	type res struct {
		tMin, tMax float64
		hit        bool
	}
	cases := []struct {
		r    Ray
		rect Rect
		want res
	}{
		// From the left
		{Ray{VecXY(-6, 2), VecXY(1, 0)}, rect, res{5, 11, true}},
		{Ray{VecXY(-6, 2), VecXY(-1, 0)}, rect, res{-11, -5, true}},
		{Ray{VecXY(-6, 2), VecXY(0, 1)}, rect, res{0, 0, false}},
		// From the right
		{Ray{VecXY(11, 2), VecXY(-1, 0)}, rect, res{6, 12, true}},
		{Ray{VecXY(11, 2), VecXY(1, 0)}, rect, res{-12, -6, true}},
		{Ray{VecXY(11, 2), VecXY(0, -1)}, rect, res{0, 0, false}},
		// Tangent
		{Ray{VecXY(5, 8), VecXY(0, -1)}, rect, res{0, 0, false}},
		// Angle
		{Ray{VecXY(-1, -5), VecXY(1, 1)}, rect, res{3 * math.Sqrt2, 6 * math.Sqrt2, true}},
	}

	for i, c := range cases {
		gotTMin, gotTMax, gotHit := c.r.IntersectRect(c.rect)
		if !gotHit && !c.want.hit {
			// ignore t values for misses
			continue
		}
		if !fEqual(gotTMin, c.want.tMin) || !fEqual(gotTMax, c.want.tMax) || gotHit != c.want.hit {
			t.Errorf("case %d: gotTMin %f, gotTMax %f, gotHit %v, want %v", i, gotTMin, gotTMax,
				gotHit, c.want)
		}
	}
}

func TestRayIntersectLine(t *testing.T) {
	type res struct {
		t   float64
		hit bool
	}
	cases := []struct {
		r      Ray
		v1, v2 Vec
		want   res
	}{
		// Hit
		{Ray{VecXY(-4, -2), VecXY(1, 0)}, VecXY(-1, -3), VecXY(9, 2), res{5, true}},
		{Ray{VecXY(-4, -2), VecXY(-1, 0)}, VecXY(-1, -3), VecXY(9, 2), res{-5, true}},
		// Miss
		{Ray{VecXY(6, 4), VecXY(1, 0)}, VecXY(-1, -3), VecXY(9, 2), res{7, false}},
		{Ray{VecXY(6, 4), VecXY(-1, 0)}, VecXY(-1, -3), VecXY(9, 2), res{-7, false}},
		// Perpendicular
		{Ray{VecXY(1, 3), VecXY(1, -2)}, VecXY(-1, -3), VecXY(9, 2), res{math.Hypot(2, 4), true}},
		{Ray{VecXY(1, 3), VecXY(-1, 2)}, VecXY(-1, -3), VecXY(9, 2), res{-math.Hypot(2, 4), true}},
		// Parallel
		{Ray{VecXY(8, -2), VecXY(2, 1)}, VecXY(-1, -3), VecXY(9, 2), res{math.Inf(-1), false}},
		{Ray{VecXY(8, -2), VecXY(-2, -1)}, VecXY(-1, -3), VecXY(9, 2), res{math.Inf(-1), false}},
		{Ray{VecXY(-8, 2), VecXY(2, 1)}, VecXY(-1, -3), VecXY(9, 2), res{math.Inf(1), false}},
		{Ray{VecXY(-8, 2), VecXY(-2, -1)}, VecXY(-1, -3), VecXY(9, 2), res{math.Inf(1), false}},

		{Ray{VecXY(-1, -1), VecXY(1, -1)}, VecXY(-1, 1), VecXY(1, -1), res{math.Inf(-1), false}},
		{Ray{VecXY(-1, -1), VecXY(-1, 1)}, VecXY(-1, 1), VecXY(1, -1), res{math.Inf(-1), false}},
		{Ray{VecXY(1, 1), VecXY(1, -1)}, VecXY(-1, 1), VecXY(1, -1), res{math.Inf(1), false}},
		{Ray{VecXY(1, 1), VecXY(-1, 1)}, VecXY(-1, 1), VecXY(1, -1), res{math.Inf(1), false}},

		// Vertical line
		// Hit
		{Ray{VecXY(-5, 2), VecXY(1, -1)}, VecXY(-3, 3), VecXY(-3, -3), res{2 * math.Sqrt2, true}},
		{Ray{VecXY(-5, 2), VecXY(-1, 1)}, VecXY(-3, 3), VecXY(-3, -3), res{-2 * math.Sqrt2, true}},
		{Ray{VecXY(-1, 2), VecXY(-1, -1)}, VecXY(-3, 3), VecXY(-3, -3), res{2 * math.Sqrt2, true}},
		{Ray{VecXY(-1, 2), VecXY(1, 1)}, VecXY(-3, 3), VecXY(-3, -3), res{-2 * math.Sqrt2, true}},
		// Miss
		{Ray{VecXY(-3, 5), VecXY(1, 0)}, VecXY(-3, 3), VecXY(-3, -3), res{0, false}},
		{Ray{VecXY(-2, 6), VecXY(-1, -1)}, VecXY(-3, 3), VecXY(-3, -3), res{math.Sqrt2, false}},
		// Perpendicular
		{Ray{VecXY(3, -2), VecXY(-1, 0)}, VecXY(-3, 3), VecXY(-3, -3), res{6, true}},
		{Ray{VecXY(3, -2), VecXY(1, 0)}, VecXY(-3, 3), VecXY(-3, -3), res{-6, true}},
		// Parallel
		{Ray{VecXY(-5, -2), VecXY(0, 1)}, VecXY(-3, 3), VecXY(-3, -3), res{math.Inf(1), false}},
		{Ray{VecXY(-5, -2), VecXY(0, -1)}, VecXY(-3, 3), VecXY(-3, -3), res{math.Inf(-1), false}},
		{Ray{VecXY(-1, -4), VecXY(0, 1)}, VecXY(-3, 3), VecXY(-3, -3), res{math.Inf(-1), false}},
		{Ray{VecXY(-1, -4), VecXY(0, -1)}, VecXY(-3, 3), VecXY(-3, -3), res{math.Inf(1), false}},

		// Horizontal line
		{Ray{VecXY(0, -2), VecXY(1, 0)}, VecXY(-3, 0), VecXY(3, 0), res{math.Inf(-1), false}},
		{Ray{VecXY(0, -2), VecXY(-1, 0)}, VecXY(-3, 0), VecXY(3, 0), res{math.Inf(1), false}},
		{Ray{VecXY(0, 2), VecXY(1, 0)}, VecXY(-3, 0), VecXY(3, 0), res{math.Inf(1), false}},
		{Ray{VecXY(0, 2), VecXY(-1, 0)}, VecXY(-3, 0), VecXY(3, 0), res{math.Inf(-1), false}},
	}

	for i, c := range cases {
		gotT, gotHit := c.r.IntersectLine(c.v1, c.v2)
		if !fEqual(gotT, c.want.t) || gotHit != c.want.hit {
			t.Errorf("case %d: gotT %f, gotHit %v, want %v", i, gotT, gotHit, c.want)
		}
	}
}
