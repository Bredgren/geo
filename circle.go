package geo

import (
	"fmt"
	"math"
)

// Circle is a 2-D circle with position (X, Y) and radius R.
type Circle struct {
	X, Y, R float64
}

// Equals returns true if the two circles position and radius are with in error e.
func (c Circle) Equals(other Circle, e float64) bool {
	return c.Pos().Equals(other.Pos(), e) && math.Abs(c.R-other.R) < e
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%g, %g, r%g)", c.X, c.Y, c.R)
}

// CircleXYR creates a Circle at position (x, y) with radius r.
func CircleXYR(x, y, r float64) Circle {
	return Circle{X: x, Y: y, R: r}
}

// CircleXY creates a unit Circle at position (x, y).
func CircleXY(x, y float64) Circle {
	return CircleXYR(x, y, 1)
}

// CircleVecR creates a Circle at position pos with radius r.
func CircleVecR(pos Vec, r float64) Circle {
	return CircleXYR(pos.X, pos.Y, r)
}

// CircleVec creates a unit Circle at position pos.
func CircleVec(pos Vec) Circle {
	return CircleVecR(pos, 1)
}

// CircleR creates a Circle centered at the origin with radius r.
func CircleR(r float64) Circle {
	return Circle{R: r}
}

// XYR returns the Circle's x, y and r values.
func (c Circle) XYR() (x, y, r float64) {
	return c.X, c.Y, c.R
}

// Pos returns the Circle's position as a Vec.
func (c Circle) Pos() Vec {
	return Vec{X: c.X, Y: c.Y}
}

// SetPos sets moves the Circel's center position.
func (c *Circle) SetPos(pos Vec) {
	c.X, c.Y = pos.XY()
}

// PosR returns the Circle's position as a Vec and radius.
func (c Circle) PosR() (pos Vec, r float64) {
	return Vec{X: c.X, Y: c.Y}, c.R
}

// Area returns the area of the Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

// Circumference returns the circumference of the Circle.
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.R
}

// SetCircumference modifies the Circle so that it's cercumference is equal to circ.
func (c *Circle) SetCircumference(circ float64) {
	c.R = circ / 2 / math.Pi
}

// Diameter returns the diameter of the Circle.
func (c Circle) Diameter() float64 {
	return 2 * c.R
}

// SetDiameter modies the Circle so that it's diameter is equal to d.
func (c *Circle) SetDiameter(d float64) {
	c.R = d / 2
}

// Top returns the y position that is tangent to the top of the Circle.
func (c Circle) Top() float64 {
	return c.Y - c.R
}

// SetTop moves the Circle so that its top side is tangent to y.
func (c *Circle) SetTop(y float64) {
	c.Y = y + c.R
}

// Left returns the x position that is tangent to the left side of the Circle.
func (c Circle) Left() float64 {
	return c.X - c.R
}

// SetLeft moves the Circle so that its left side is tangent to x.
func (c *Circle) SetLeft(x float64) {
	c.X = x + c.R
}

// Right returns the x position that is tangent to the right side of the Circle.
func (c Circle) Right() float64 {
	return c.X + c.R
}

// SetRight moves the Circle so that its right side is tangent to x.
func (c *Circle) SetRight(x float64) {
	c.X = x - c.R
}

// Bottom returns the y position that is tangent to the bottom of the Circle.
func (c Circle) Bottom() float64 {
	return c.Y + c.R
}

// SetBottom moves the Circle so that its bottom side is tangent to y.
func (c *Circle) SetBottom(y float64) {
	c.Y = y - c.R
}

// TopLeft returns the coordinates of the top left corner of the Circle's bounding box.
func (c Circle) TopLeft() (x, y float64) {
	return c.Left(), c.Top()
}

// SetTopLeft moves the Circle so that the top left corner of its bounding box is at the
// given coordinates.
func (c *Circle) SetTopLeft(x, y float64) {
	c.SetLeft(x)
	c.SetTop(y)
}

// TopMid returns the coordinates of the top of the Circle above its center.
func (c Circle) TopMid() (x, y float64) {
	return c.X, c.Top()
}

// SetTopMid moves the Circle so that the top mid position is at (x, y).
func (c *Circle) SetTopMid(x, y float64) {
	c.X = x
	c.SetTop(y)
}

// TopRight returns the position of the top right corner of the Circle's bounding box.
func (c Circle) TopRight() (x, y float64) {
	return c.Right(), c.Top()
}

// SetTopRight moves the Circle so that the top right corner of its bounding box is at
// position (x, y).
func (c *Circle) SetTopRight(x, y float64) {
	c.SetRight(x)
	c.SetTop(y)
}

// LeftMid returns the position of the left of the Circle, inline with its center.
func (c Circle) LeftMid() (x, y float64) {
	return c.Left(), c.Y
}

// SetLeftMid moves the Circle so that the left mid position is at (x, y).
func (c *Circle) SetLeftMid(x, y float64) {
	c.SetLeft(x)
	c.Y = y
}

// Mid returns the Circle's center x and y position.
func (c Circle) Mid() (x, y float64) {
	return c.X, c.Y
}

// SetMid moves the Circle so that its center is at position (x, y).
func (c *Circle) SetMid(x, y float64) {
	c.X, c.Y = x, y
}

// RightMid returns the position of the right of the Circle, inline with its center.
func (c Circle) RightMid() (x, y float64) {
	return c.Right(), c.Y
}

// SetRightMid moves the Circle so that the right mid position is at (x, y).
func (c *Circle) SetRightMid(x, y float64) {
	c.SetRight(x)
	c.Y = y
}

// BottomLeft returns the position of the bottom left corner of the Circle's bounding box.
func (c Circle) BottomLeft() (x, y float64) {
	return c.Left(), c.Bottom()
}

// SetBottomLeft moves the Circle so that the bottom left corner of its bounding box is
// at position (x, y).
func (c *Circle) SetBottomLeft(x, y float64) {
	c.SetLeft(x)
	c.SetBottom(y)
}

// BottomMid returns the position of the bottom of the Circle, bellow its center.
func (c Circle) BottomMid() (x, y float64) {
	return c.X, c.Bottom()
}

// SetBottomMid moves the Circle so the bottom mod position is at (x, y).
func (c *Circle) SetBottomMid(x, y float64) {
	c.X = x
	c.SetBottom(y)
}

// BottomRight returns the position of the bottom right corner of the Circle's bounding box.
func (c Circle) BottomRight() (x, y float64) {
	return c.Right(), c.Bottom()
}

// SetBottomRight moves the Circle so that the bottom right corner of its bounding box is at
// position (x, y).
func (c *Circle) SetBottomRight(x, y float64) {
	c.SetRight(x)
	c.SetBottom(y)
}

// Inflate keeps the same center but changes the size by the given amount.
func (c *Circle) Inflate(dr float64) {
	c.R += dr
}

// Inflated returns a new Circle with the same center but changes the size by the given amount.
func (c Circle) Inflated(dr float64) Circle {
	c.R += dr
	return c
}

// Move moves the center of the Circle by the given amount.
func (c *Circle) Move(dx, dy float64) {
	c.X += dx
	c.Y += dy
}

// Moved returns a new Circle with the same size but with the center moved by the given amount.
func (c Circle) Moved(dx, dy float64) Circle {
	c.X += dx
	c.Y += dy
	return c
}

// Normalize modifies the radius to be positive if needed.
func (c *Circle) Normalize() {
	c.R = math.Abs(c.R)
}

// Normalized returns a new Circle with a positive radius.
func (c Circle) Normalized() Circle {
	c.R = math.Abs(c.R)
	return c
}

// Contains returns true if the other Circle is completely inside this one.
func (c Circle) Contains(other Circle) bool {
	return c.R >= other.R && c.Pos().Dist2(other.Pos()) < (c.R-other.R)*(c.R-other.R)
}

// CollideCircle returns true if other overlaps with this one.
func (c Circle) CollideCircle(other Circle) bool {
	return c.Pos().Dist2(other.Pos()) < (c.R+other.R)*(c.R+other.R)
}

// CollidePoint returns true if the given point is inside the Circle.
func (c Circle) CollidePoint(x, y float64) bool {
	return c.Pos().Dist2(VecXY(x, y)) < c.R*c.R
}

// CollideCircleList returns the index of the first Circle this one collides with. If there
// is no collision then ok is false and i is undefined.
func (c Circle) CollideCircleList(others []Circle) (i int, ok bool) {
	for i, other := range others {
		if c.CollideCircle(other) {
			return i, true
		}
	}
	return
}

// CollideCircleListAll returns a list of indices of the Circles that collide with this one,
// or an empty list if none.
func (c Circle) CollideCircleListAll(others []Circle) []int {
	list := make([]int, 0, len(others))
	for i, other := range others {
		if c.CollideCircle(other) {
			list = append(list, i)
		}
	}
	return list
}

// Union modifies this Circle to be the smallest one that would contain both itself and
// the given one
func (c *Circle) Union(other Circle) {
	v1 := c.Pos()
	v2 := other.Pos()
	if v1 == v2 {
		return
	}
	v12 := v2.Minus(v1).Normalized()
	p1 := v1.Plus(v12.Times(-c.R))
	p2 := v2.Plus(v12.Times(other.R))
	longVec := p2.Minus(p1)
	midPoint := p1.Plus(longVec.DividedBy(2))
	c.X, c.Y = midPoint.X, midPoint.Y
	c.R = longVec.Len() / 2
}

// Unioned returns the smallest Circle that contains both this one and other.
func (c Circle) Unioned(other Circle) Circle {
	c.Union(other)
	return c
}

// // EnclosingCircle returns the smallest Circle that encloses all the given points.
// func EnclosingCircle(points []Vec) Circle {
// 	shuffled := make([]Vec, len(points))
// 	perm := rand.Perm(len(points))
// 	for i, v := range perm {
// 		shuffled[v] = points[i]
// 	}
// 	return enclosingCircle(shuffled, []Vec{})
// }
//
// func enclosingCircle(points []Vec, pointsOnBoundary []Vec) Circle {
// 	// http://www.sunshine2k.de/coding/java/Welzl/Welzl.html
// 	// if (P is empty or |R| = 3) then
// 	//        D := calcDiskDirectly(R)
// 	//   else
// 	//       choose a p from P randomly;
// 	//       D := sed(P - {p}, R);
// 	//       if (p lies NOT inside D) then
// 	//           D := sed(P - {p}, R u {p});
// 	//   return D;
// 	return Circle{}
// }
//
// // CircleUnion retusn the smallest Circle that encloses all the given Circles.
// func CircleUnion(circles []Circle) Circle {
// 	// https://bl.ocks.org/mbostock/29c534ff0b270054a01c
// 	return Circle{}
// }

// BoundingRect returns the smallest Rect that surrounds the Circle.
func (c Circle) BoundingRect() Rect {
	size := c.Diameter()
	return RectXYWH(c.Left(), c.Top(), size, size)
}

// InscribedRect returns the largest square that fits inside the Circle.
func (c Circle) InscribedRect() Rect {
	size := math.Sqrt(2) * c.R
	r := RectWH(size, size)
	r.SetMid(c.X, c.Y)
	return r
}

// CircleCircumscribe returns the smallest circle that surrounds the Rect.
func CircleCircumscribe(r Rect) Circle {
	diameter := VecXY(r.TopLeft()).Minus(VecXY(r.BottomRight())).Len()
	return CircleXYR(r.MidX(), r.MidY(), diameter/2)
}

// CircleInscribe returns the largest Circle that fits insided the Rect, and shares its
// center with the Rect.
func CircleInscribe(r Rect) Circle {
	r.Normalized()
	radius := math.Min(r.W, r.H) / 2
	return CircleXYR(r.MidX(), r.MidY(), radius)
}

// PointAt returns the position of the point along the Circle's edge that is at the given
// angle (from +x axis).
func (c Circle) PointAt(radians float64) (x, y float64) {
	x = c.X + math.Cos(radians)*c.R
	y = c.Y - math.Sin(radians)*c.R
	return
}

// CollideRect returns true if the Circle is colliding with the Rect.
func (c Circle) CollideRect(r Rect) bool {
	// Check if the closest point to the circle in the Rect is within the circle
	return c.CollidePoint(c.Pos().Clamped(r).XY())
}

// CollideRectList returns the index of the first Rect the Circle collides with. If there
// is no collision then ok is false and i is undefined.
func (c Circle) CollideRectList(rs []Rect) (i int, ok bool) {
	for i, r := range rs {
		if c.CollideRect(r) {
			return i, true
		}
	}
	return
}

// CollideRectListAll returns a list of indices of the Rects that collide with the Circle, or an
// empty list if none.
func (c Circle) CollideRectListAll(rs []Rect) []int {
	list := make([]int, 0, len(rs))
	for i, r := range rs {
		if c.CollideRect(r) {
			list = append(list, i)
		}
	}
	return list
}
