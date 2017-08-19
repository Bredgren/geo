package geo

import (
	"fmt"
	"math"
)

// Ray is a 2-D ray with an origin and direction. The direction Vec does not have to be
// normalized, the methods of Ray will use the normalized direction when needed.
type Ray struct {
	Origin, Direction Vec
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(%s, %s)", r.Origin, r.Direction)
}

// RayAngle creates a new Ray whose direction is angled from the +x axis by the given radians.
func RayAngle(origin Vec, radians float64) Ray {
	dir := Vec{X: 1}
	dir.Rotate(radians)
	return Ray{Origin: origin, Direction: dir}
}

// At returns the Vec that distance t from the Ray's origin in the Ray's direction. If
// the Ray's direction has length 0 then the origin will always be returned.
func (r Ray) At(t float64) Vec {
	return r.Origin.Plus(r.Direction.WithLen(t))
}

// IntersectCircle tests whether the Ray intersects the Circle. It returns both intersection
// points as tMin and tMax (tMin <= tMax). The t values can be negative, in which case
// the intersection point is behind the Ray (if tMin < 0 and tMax > 0 then the point is
// in the Circle). The function will return true for hit if the Ray intersects either
// behind or in front of the Ray. If hit is false then tMin and tMax are undefined. Use
// Ray.At with tMin or tMax to get the actual intersection points as a Vec.
func (r Ray) IntersectCircle(c Circle) (tMin, tMax float64, hit bool) {
	center := c.Pos()
	toCircle := center.Minus(r.Origin)
	t := toCircle.Dot(r.Direction.Normalized())
	nearestToCircle := r.At(t)

	dist2 := nearestToCircle.Dist2(center)
	r2 := c.R * c.R

	if dist2 >= r2 {
		return
	}

	dt := math.Sqrt(r2 - dist2)
	return t - dt, t + dt, true
}

// IntersectRect tests whether the Ray intersects the Rect. It returns both intersection
// points as tMin and tMax (tMin <= tMax). The t values can be negative, in which case
// the intersection point is behind the Ray (if tMin < 0 and tMax > 0 the nthe point is
// in the Rect). The function will return true for hit if the Ray intersects either behind
// or in front oe the Ray. If hit is false then tMin and tMax are undefined. Use Ray.At
// with tMin or tMax to get the actual intersection points as a Vec.
func (r Ray) IntersectRect(rect Rect) (tMin, tMax float64, hit bool) {
	// https://tavianator.com/fast-branchless-raybounding-box-intersections/
	dir := r.Direction.Normalized()

	tLeft := (rect.Left() - r.Origin.X) / dir.X
	tRight := (rect.Right() - r.Origin.X) / dir.X

	tMin = math.Min(tLeft, tRight)
	tMax = math.Max(tLeft, tRight)

	tTop := (rect.Top() - r.Origin.Y) / dir.Y
	tBottom := (rect.Bottom() - r.Origin.Y) / dir.Y

	tMin = math.Max(tMin, math.Min(tTop, tBottom))
	tMax = math.Min(tMax, math.Max(tTop, tBottom))

	return tMin, tMax, tMax >= tMin
}

// IntersectLine tests whether the Ray intersects the line between v1 and v2. The return
// value of t is how far the Ray must be extended to hit the line, and will be negative
// if it hits behind the Ray. The value of hit will be true if the intersection point is
// in between v1 and v2. However the value of t will still be accurate if the Ray intersects
// outside of those points. If the Ray is parallel to the line between v1 and v2 then
// t will +/-Inf, no promises are made about its sign but math.IsInf(t, 0) will be true.
func (r Ray) IntersectLine(v1, v2 Vec) (t float64, hit bool) {
	// https://rootllama.wordpress.com/2014/06/20/ray-line-segment-intersection-test-in-2d/
	dir := r.Direction.Normalized()
	a := r.Origin.Minus(v1)
	b := v2.Minus(v1)
	c := VecXY(-dir.Y, dir.X)
	t1 := b.Cross(a) / b.Dot(c)
	t2 := a.Dot(c) / b.Dot(c)
	return t1, 0 <= t2 && t2 <= 1
}
