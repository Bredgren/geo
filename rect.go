package geo

import "math"

// Rect defines rectangular coordinates, by the top left corner (X, Y), width W and height H.
type Rect struct {
	X, Y, W, H float64
}

// RectCorners creates a rect given top lect and bottom right coners.
func RectCorners(x1, y1, x2, y2 float64) Rect {
	return Rect{X: x1, Y: y1, W: x2 - x1, H: y2 - y1}
}

// RectCornersVec creates a rect given top lect and bottom right coners.
func RectCornersVec(topLeft, bottomRight Vec) Rect {
	return Rect{X: topLeft.X, Y: topLeft.Y, W: bottomRight.X - topLeft.X, H: bottomRight.Y - topLeft.Y}
}

// Top returns the top boundary.
func (r Rect) Top() (top float64) {
	return r.Y
}

// SetTop sets the top boundary.
func (r *Rect) SetTop(top float64) {
	r.Y = top
}

// Bottom returns the bottom boundary.
func (r Rect) Bottom() float64 {
	return r.Y + r.H
}

// SetBottom sets the bottom boundary.
func (r *Rect) SetBottom(bottom float64) {
	r.Y = bottom - r.H
}

// Left returns the left boundary.
func (r Rect) Left() float64 {
	return r.X
}

// SetLeft sets the left boundary.
func (r *Rect) SetLeft(left float64) {
	r.X = left
}

// Right returns the right boundary.
func (r Rect) Right() float64 {
	return r.X + r.W
}

// SetRight sets the right boundary.
func (r *Rect) SetRight(right float64) {
	r.X = right - r.W
}

// Size returns the width and height.
func (r Rect) Size() (w, h float64) {
	return r.W, r.H
}

// SetSize sets the width and height.
func (r *Rect) SetSize(w, h float64) {
	r.W = w
	r.H = h
}

// TopLeft returns the coordinates of the top left corner.
func (r Rect) TopLeft() (x, y float64) {
	return r.Left(), r.Top()
}

// SetTopLeft sets the coordinates of the top left corner.
func (r *Rect) SetTopLeft(x, y float64) {
	r.SetLeft(x)
	r.SetTop(y)
}

// BottomLeft returns the coordinates of the bottom left corner.
func (r Rect) BottomLeft() (x, y float64) {
	return r.Left(), r.Bottom()
}

// SetBottomLeft set the coordinates of the bottom left corner.
func (r *Rect) SetBottomLeft(x, y float64) {
	r.SetLeft(x)
	r.SetBottom(y)
}

// TopRight returns the coordinates of the top right corner.
func (r Rect) TopRight() (x, y float64) {
	return r.Right(), r.Top()
}

// SetTopRight sets the coordinates of the top right corner.
func (r *Rect) SetTopRight(x, y float64) {
	r.SetRight(x)
	r.SetTop(y)
}

// BottomRight returns the coordinates of the bottom right corner.
func (r Rect) BottomRight() (x, y float64) {
	return r.Right(), r.Bottom()
}

// SetBottomRight sets the coordinates of the bottom right corner.
func (r *Rect) SetBottomRight(x, y float64) {
	r.SetRight(x)
	r.SetBottom(y)
}

// TopMid returns the coordinates at the top of the rectangle above the center.
func (r Rect) TopMid() (x, y float64) {
	return r.MidX(), r.Top()
}

// SetTopMid sets the coordinates at the top of the rectangle above the center.
func (r *Rect) SetTopMid(x, y float64) {
	r.SetMidX(x)
	r.SetTop(y)
}

// BottomMid returns the coordinates at the bottom of the rectangle below the center.
func (r Rect) BottomMid() (x, y float64) {
	return r.MidX(), r.Bottom()
}

// SetBottomMid sets the coordinates at the bottom of the rectangle below the center.
func (r *Rect) SetBottomMid(x, y float64) {
	r.SetMidX(x)
	r.SetBottom(y)
}

// LeftMid returns the coordinates at the left of the rectangle in line with the center.
func (r Rect) LeftMid() (x, y float64) {
	return r.Left(), r.MidY()
}

// SetLeftMid sets the coordinates at the left of the rectangle in line with the center.
func (r *Rect) SetLeftMid(x, y float64) {
	r.SetLeft(x)
	r.SetMidY(y)
}

// RightMid returns the coordinates at the right of the rectangle in line with the center.
func (r Rect) RightMid() (x, y float64) {
	return r.Right(), r.MidY()
}

// SetRightMid sets the coordinates at the right of the rectangle in line with the center.
func (r *Rect) SetRightMid(x, y float64) {
	r.SetRight(x)
	r.SetMidY(y)
}

// Mid returns the center coordinates.
func (r Rect) Mid() (x, y float64) {
	return r.MidX(), r.MidY()
}

// SetMid sets the center coordinates.
func (r *Rect) SetMid(x, y float64) {
	r.SetMidX(x)
	r.SetMidY(y)
}

// MidX returns the center x coordinates
func (r Rect) MidX() float64 {
	return r.X + r.W/2
}

// SetCenterX sets the center x coordinates.
func (r *Rect) SetMidX(x float64) {
	r.X = x - r.W/2
}

// CenterY returns the center y coordinates.
func (r Rect) MidY() float64 {
	return r.Y + r.H/2
}

// SetCenterY set the center y coordinates.
func (r *Rect) SetMidY(y float64) {
	r.Y = y - r.H/2
}

// Area returns the area of the rectangle.
func (r Rect) Area() float64 {
	return r.W * r.H
}

// Move moves the Rect by the given offset, in place.
func (r *Rect) Move(dx, dy float64) {
	r.X += dx
	r.Y += dy
}

// Moved returns a new Rect moved by the given offset relative to this one.
func (r Rect) Moved(dx, dy float64) Rect {
	return Rect{X: r.X + dx, Y: r.Y + dy, W: r.W, H: r.H}
}

// Inflate keeps the same center but changes the size by the given amount, in place.
func (r *Rect) Inflate(dw, dh float64) {
	r.X -= dw / 2
	r.Y -= dh / 2
	r.W += dw
	r.H += dh
}

// Inflated returns a new Rect with the same center whose size is chaged by the given amount.
func (r Rect) Inflated(dw, dh float64) Rect {
	return Rect{X: r.X - dw/2, Y: r.Y - dh/2, W: r.W + dw, H: r.H + dh}
}

// Clamp moves this Rect so that it is within bounds. If it is too large than it is
// centered within bounds.
func (r *Rect) Clamp(bounds Rect) {
	if r.W > bounds.W {
		r.SetMidX(bounds.MidX())
	} else {
		r.X = Clamp(r.X, bounds.X, bounds.Right()-r.W)
	}
	if r.H > bounds.H {
		r.SetMidY(bounds.MidY())
	} else {
		r.Y = Clamp(r.Y, bounds.Y, bounds.Bottom()-r.H)
	}
}

// Clamped returns a new Rect that is moved to be within bounds. If it is too large than it
// is centered within bounds.
func (r Rect) Clamped(bounds Rect) Rect {
	var newX, newY float64
	if r.W > bounds.W {
		newX = bounds.MidX() - r.W/2
	} else {
		newX = Clamp(r.X, bounds.X, bounds.Right()-r.W)
	}
	if r.H > bounds.H {
		newY = bounds.MidY() - r.H/2
	} else {
		newY = Clamp(r.Y, bounds.Y, bounds.Bottom()-r.H)
	}
	return Rect{X: newX, Y: newY, W: r.W, H: r.H}
}

// Intersect returns a new Rect that marks the area where the two overlap. If there is
// no intersection the returned Rect will have 0 size.
func (r Rect) Intersect(other Rect) Rect {
	newX := math.Max(r.X, other.X)
	newY := math.Max(r.Y, other.Y)
	return Rect{
		X: newX,
		Y: newY,
		W: math.Min(r.Right(), other.Right()) - newX,
		H: math.Min(r.Bottom(), other.Bottom()) - newY,
	}
}

// Union modifies this Rect to contain both itself and other.
func (r *Rect) Union(other Rect) {
	newX := math.Min(r.X, other.X)
	newY := math.Min(r.Y, other.Y)
	r.W = math.Max(r.Right(), other.Right()) - newX
	r.H = math.Max(r.Bottom(), other.Bottom()) - newY
	r.X = newX
	r.Y = newY
}

// Unioned retruns a new Rect that contain both Rects.
func (r Rect) Unioned(other Rect) Rect {
	newX := math.Min(r.X, other.X)
	newY := math.Min(r.Y, other.Y)
	r.W = math.Max(r.Right(), other.Right()) - newX
	r.H = math.Max(r.Bottom(), other.Bottom()) - newY
	r.X = newX
	r.Y = newY
	return r
}

// Fit modifies this Rect so that it is moved and resized to fit within bounds while maintaining
// its original aspect ratio.
func (r *Rect) Fit(bounds Rect) {
	newW := bounds.H * (r.W / r.H)
	if newW <= bounds.W {
		r.X = Clamp(r.X, bounds.X, bounds.Right()-newW)
		r.Y = bounds.Y
		r.W = newW
		r.H = bounds.H
		return
	}
	newH := bounds.W * (r.H / r.W)
	r.X = bounds.X
	r.Y = Clamp(r.Y, bounds.Y, bounds.Bottom()-newH)
	r.W = bounds.W
	r.H = newH
}

// Fitted returns a new Rect that is moved and resized to fit within bounds while maintaining
// its original aspect ratio.
func (r Rect) Fitted(bounds Rect) Rect {
	newW := bounds.H * (r.W / r.H)
	if newW <= bounds.W {
		return Rect{X: Clamp(r.X, bounds.X, bounds.Right()-newW), Y: bounds.Y, W: newW, H: bounds.H}
	}
	newH := bounds.W * (r.H / r.W)
	return Rect{X: bounds.X, Y: Clamp(r.Y, bounds.Y, bounds.Bottom()-newH), W: bounds.W, H: newH}
}

// Normalize fips the Rect in place if its size is negative.
func (r *Rect) Normalize() {
	if r.W < 0 {
		r.X += r.W
		r.W = -r.W
	}
	if r.H < 0 {
		r.Y += r.H
		r.H = -r.H
	}
}

// Normalized returns a fipped but equivalent Rect if its size is negative.
func (r Rect) Normalized() Rect {
	r.Normalize()
	return r
}

// Contains returns true if other is completely inside this one.
func (r Rect) Contains(other Rect) bool {
	return other.X >= r.X && other.Y >= r.Y && other.Right() <= r.Right() && other.Bottom() <= r.Bottom()
}

// CollidePoint returns true if the point is within the Rect. A point along the right or
// bottom edge is not considered inside.
func (r Rect) CollidePoint(x, y float64) bool {
	return x >= r.X && x < r.Right() && y >= r.Y && y < r.Bottom()
}

// CollideRect returns true if the Rects overlap.
func (r Rect) CollideRect(other Rect) bool {
	return r.X < other.Right() && r.Right() > other.X && r.Y < other.Bottom() && r.Bottom() > other.Y
}

// CollideRectList returns the index of the first Rect this one collides with. If it there
// is no collision then ok is false and i is undefined.
func (r Rect) CollideRectList(others []Rect) (i int, ok bool) {
	for i, other := range others {
		if r.CollideRect(other) {
			return i, ok
		}
	}
	return
}

// CollideRectListAll returns a list of indices of the Rects that collide with this one, or an
// empty list if none.
func (r Rect) CollideRectListAll(others []Rect) []int {
	list := make([]int, 0, len(others))
	for i, other := range others {
		if r.CollideRect(other) {
			list = append(list, i)
		}
	}
	return list
}

// RectUnion returns a Rect that contains all the given Rects. An empty list returns a
// Rect with size 0.
func RectUnion(rects []Rect) Rect {
	if len(rects) == 0 {
		return Rect{}
	}

	newX, newY := rects[0].X, rects[0].Y
	for _, r := range rects {
		newX = math.Min(newX, r.X)
		newY = math.Min(newY, r.Y)
	}

	farRight, farBottom := rects[0].Right(), rects[0].Bottom()
	for _, r := range rects {
		farRight = math.Max(farRight, r.Right())
		farBottom = math.Max(farBottom, r.Bottom())
	}

	return Rect{X: newX, Y: newY, W: farRight - newX, H: farBottom - newY}
}
