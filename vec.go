package geo

import (
	"fmt"
	"image"
	"math"
	"math/rand"
)

// Vec is a 2-D vector. Many of the functions for Vec have two versions, one that modifies
// the Vec and one that returns a new Vec. Their names follow a convention that is hopefully
// intuitive. For example, when working with Vec as a value you use v1.Plus(v2) which returns
// a new value and reads like when working with other value types such as "1 + 2". The other
// function, v1.Add(v2), adds v2 to v1 which actually modifies v1.
type Vec struct {
	X, Y float64
}

// Vec0 is the zero vector.
var Vec0 = Vec{}

// VecXY constructs a Vec from its arguments. Useful for making a Vec from a function that
// returns two floats.
func VecXY(x, y float64) Vec {
	return Vec{X: x, Y: y}
}

// VecXYi creates a Vec from two ints.
func VecXYi(x, y int) Vec {
	return Vec{X: float64(x), Y: float64(y)}
}

// VecPoint creates a Vec from an image.Point.
func VecPoint(p image.Point) Vec {
	return Vec{X: float64(p.X), Y: float64(p.Y)}
}

// VecLA creates a Vec from a length and an angle.
func VecLA(length, rad float64) Vec {
	v := Vec{X: length}
	v.Rotate(rad)
	return v
}

// Point converts a Vec into image.Point, discarding any fractional components to X and Y.
func (v Vec) Point() image.Point {
	return image.Pt(int(v.X), int(v.Y))
}

func (v Vec) String() string {
	return fmt.Sprintf("Vec(%g, %g)", v.X, v.Y)
}

// XY returns the Vec's components. Useful for passing a Vec to a function that takes
// x and y individually.
func (v Vec) XY() (x, y float64) {
	return v.X, v.Y
}

// Set modifies the Vec's x and Y to those given.
func (v *Vec) Set(x, y float64) {
	v.X, v.Y = x, y
}

// Equals returns true if the corresponding components of the vectors are within the error e.
func (v Vec) Equals(v2 Vec, e float64) bool {
	return math.Abs(v.X-v2.X) < e && math.Abs(v.Y-v2.Y) < e
}

// Len returns the length of the vector.
func (v Vec) Len() float64 {
	return math.Hypot(v.X, v.Y)
}

// Len2 is the length of the vector squared.
func (v Vec) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}

// SetLen sets the length of the vector. Negative lengths will flip the vectors direction.
// If v's length is zero then it will remain unchanged.
func (v *Vec) SetLen(l float64) {
	len := v.Len()
	if len == 0 {
		return
	}
	v.X = v.X / len * l
	v.Y = v.Y / len * l
}

// WithLen returns a new vector in the same direction as v but with the given length.
// It v's length is 0 then the zero vector is returned.
func (v Vec) WithLen(l float64) Vec {
	len := v.Len()
	if len == 0 {
		return Vec{}
	}
	v.X = v.X / len * l
	v.Y = v.Y / len * l
	return v
}

// Dist returns the distance between the two vectors.
func (v Vec) Dist(v2 Vec) float64 {
	return math.Hypot(v.X-v2.X, v.Y-v2.Y)
}

// Dist2 returns the distance squared between the two vectors.
func (v Vec) Dist2(v2 Vec) float64 {
	return (v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y)
}

// Add modifies v to be the sum of v2 and itself.
func (v *Vec) Add(v2 Vec) {
	v.X += v2.X
	v.Y += v2.Y
}

// Plus returns a new vector that is the sum of the two vectors.
func (v Vec) Plus(v2 Vec) Vec {
	return Vec{X: v.X + v2.X, Y: v.Y + v2.Y}
}

// Sub modifies v to be the difference between itself and v2.
func (v *Vec) Sub(v2 Vec) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Minus returns a new vector that is the difference of the two vectors.
func (v Vec) Minus(v2 Vec) Vec {
	return Vec{X: v.X - v2.X, Y: v.Y - v2.Y}
}

// Mul modifies v to be itself times n.
func (v *Vec) Mul(n float64) {
	v.X *= n
	v.Y *= n
}

// Times returns a new vector that is v times n.
func (v Vec) Times(n float64) Vec {
	return Vec{X: v.X * n, Y: v.Y * n}
}

// Div modifies v to be itself divided by n.
func (v *Vec) Div(n float64) {
	v.X /= n
	v.Y /= n
}

// DividedBy returns a new vector that is v divided by n.
func (v Vec) DividedBy(n float64) Vec {
	return Vec{X: v.X / n, Y: v.Y / n}
}

// Normalize modifies v to be of length one in the same direction.
func (v *Vec) Normalize() {
	len := v.Len()
	v.X = v.X / len
	v.Y = v.Y / len
}

// Normalized returns a new vector of length one in the same direction as v.
func (v Vec) Normalized() Vec {
	len := v.Len()
	return Vec{X: v.X / len, Y: v.Y / len}
}

// Dot returns the dot product between the two vectors.
func (v Vec) Dot(v2 Vec) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Project modifies v to be the vector that is the projection of v onto v2.
func (v *Vec) Project(v2 Vec) {
	v2.Normalize()
	v2.Mul(v.X*v2.X + v.Y*v2.Y)
	v.X, v.Y = v2.X, v2.Y
}

// Projected return the Vec that is v projected onto v2.
func (v Vec) Projected(v2 Vec) Vec {
	v2.Normalize()
	v2.Mul(v.X*v2.X + v.Y*v2.Y)
	return v2
}

// Limit constrains the length of the vector be no greater than len. If the vector is already
// less than len then no change is made.
func (v *Vec) Limit(len float64) {
	l := v.Len()
	if l > len {
		v.SetLen(len)
	}
}

// Limited returns a new vector in the same direction as v with length no greater than
// len. The vector returned will be equivalent to v if v.Len() <= len.
func (v Vec) Limited(len float64) Vec {
	l := v.Len()
	if l > len {
		v.SetLen(len)
	}
	return v
}

// Angle returns the radians relative to the positive +x-axis (counterclockwise in screen
// coordinates). The returned value is in the range [-π, π).
func (v Vec) Angle() float64 {
	return -math.Atan2(v.Y, v.X)
}

// AngleFrom returns the radians from v2 to v (counterclockwise in screen coordinates).
// The returned value is in the range [-π, π).
func (v Vec) AngleFrom(v2 Vec) float64 {
	r := math.Atan2(v2.Y, v2.X) - math.Atan2(v.Y, v.X)
	if r < -math.Pi {
		r += 2 * math.Pi
	}
	if r >= math.Pi {
		r -= 2 * math.Pi
	}
	return r
}

// Rotate rotates the vector (counterclockwise in screen coordinates) by the given radians.
func (v *Vec) Rotate(rad float64) {
	v.X, v.Y = v.X*math.Cos(rad)+v.Y*math.Sin(rad), -v.X*math.Sin(rad)+v.Y*math.Cos(rad)
}

// Rotated returns a new vector that is equal to this one rotated (counterclockwise in
// screen coordinates) by the given radians.
func (v Vec) Rotated(rad float64) Vec {
	return Vec{X: v.X*math.Cos(rad) + v.Y*math.Sin(rad), Y: -v.X*math.Sin(rad) + v.Y*math.Cos(rad)}
}

// RandVec returns a unit vector in a random direction.
func RandVec() Vec {
	rad := rand.Float64() * 2 * math.Pi
	return Vec{X: math.Cos(rad), Y: math.Sin(rad)}
}

// Mod "wraps" the vector around the Rect so it stays with its bounds.
func (v *Vec) Mod(r Rect) {
	v.X = Mod(v.X-r.Left(), r.W) + r.Left()
	v.Y = Mod(v.Y-r.Top(), r.H) + r.Top()
}

// Modded is like Mod but returns a new Vec instead of modifing this one.
func (v Vec) Modded(r Rect) Vec {
	v.Mod(r)
	return v
}

// Map takes v relative to r1 and moves it to the same relative position in r2. E.g.
// if v is in the center of r1 it will be moved to the center of r2.
func (v *Vec) Map(r1, r2 Rect) {
	v.X = Map(v.X, r1.Left(), r1.Right(), r2.Left(), r2.Right())
	v.Y = Map(v.Y, r1.Top(), r1.Bottom(), r2.Top(), r2.Bottom())
}

// Mapped is like Map but returns a new Vec instead of modifing this one.
func (v Vec) Mapped(r1, r2 Rect) Vec {
	v.Map(r1, r2)
	return v
}

// Clamp modifies the Vec to be closest point within Rect to the original Vec.
func (v *Vec) Clamp(r Rect) {
	v.X = Clamp(v.X, r.Left(), r.Right())
	v.Y = Clamp(v.Y, r.Bottom(), r.Top())
}

// Clamped returns the closest point within Rect to the Vec.
func (v Vec) Clamped(r Rect) Vec {
	v.X = Clamp(v.X, r.Left(), r.Right())
	v.Y = Clamp(v.Y, r.Bottom(), r.Top())
	return v
}
