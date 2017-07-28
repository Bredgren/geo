package geo

import (
	"math"
	"math/rand"
)

// NumGen (Number Generator) is a function that returns a number.
type NumGen func() float64

// ConstNum returns a NumGen that always returns n.
func ConstNum(n float64) NumGen {
	return func() float64 {
		return n
	}
}

// RandNum returns a NumGen that returns a uniform random number in [min(a, b), max(a, b)).
func RandNum(a, b float64) NumGen {
	if a > b {
		a, b = b, a
	}
	width := b - a
	return func() float64 {
		return rand.Float64()*width + a
	}
}

// RandRadius returns a NumGen that returns a random, uniform circle radius in
// [min(radius1, radius2), max(radius1, radius2)). This function is undefined for negative
// radii.
func RandRadius(radius1, radius2 float64) NumGen {
	if radius1 > radius2 {
		radius1, radius2 = radius2, radius1
	}
	if radius1 == radius2 {
		return func() float64 {
			return radius1
		}
	}
	unitMin := radius1 / radius2
	unitMin *= unitMin
	return func() float64 {
		return math.Sqrt(rand.Float64()*(1-unitMin)+unitMin) * radius2
	}
}

// VecGen (Vector Generator) is a function that returns a vector.
type VecGen func() Vec

// StaticVec returns a VecGen that always returns the constant vector v.
func StaticVec(v Vec) VecGen {
	return func() Vec {
		return v
	}
}

// DynamicVec returns a VecGen that always returns a copy of the vector pointed to by v.
func DynamicVec(v *Vec) VecGen {
	return func() Vec {
		return *v
	}
}

// OffsetVec returns a VecGen that adds offset to gen. For example, if you wanted to use
// RandCircle as an initial position you might use the following to center the circle
// at position 100, 100.
//  OffsetVec(RandVecCircle(5, 10), StaticVec(Vec{X: 100, Y: 100}))
func OffsetVec(gen VecGen, offset VecGen) VecGen {
	return func() Vec {
		return gen().Plus(offset())
	}
}

// RandVecCircle returns a VecGen that will generate a random vector whose length is
// in [min(radius1, radius2), max(radius1, radius2)), and is uniformly distributed within
// the circle.
func RandVecCircle(radius1, radius2 float64) VecGen {
	if radius1 > radius2 {
		radius1, radius2 = radius2, radius1
	}
	return func() Vec {
		return RandVec().Times(circleRadius(radius1, radius2))
	}
}

// RandVecArc returns a VecGen that will generate a random vector within the slice of a
// circle defined by the parameters. The radians are relative to the +x axis. The length
// of the vector will be within [min(radius1, radius2), max(radius1, radius2)) and the
// angle will be within [min(radians1, radians2), max(radians1, radians2)).
func RandVecArc(radius1, radius2, radians1, radians2 float64) VecGen {
	if radius1 > radius2 {
		radius1, radius2 = radius2, radius1
	}
	if radians1 > radians2 {
		radians1, radians2 = radians2, radians1
	}
	return func() Vec {
		r := circleRadius(radius1, radius2)
		rad := rand.Float64()*(radians2-radians1) + radians1
		return Vec{X: r}.Rotated(rad)
	}
}

// RandVecRect returns a VecGen that will generate a random vector within the given Rect.
func RandVecRect(rect Rect) VecGen {
	return func() Vec {
		return Vec{
			X: rand.Float64()*rect.W + rect.X,
			Y: rand.Float64()*rect.H + rect.Y,
		}
	}
}

// RandVecRects returns a VecGen that will generate a random vector that is uniformly
// distributed between all the given rects. If the slice given is empty then the zero
// vector is returned.
func RandVecRects(rects []Rect) VecGen {
	if len(rects) == 0 {
		return func() Vec { return Vec{} }
	}
	areas := make([]float64, len(rects))
	for i := range rects {
		areas[i] = rects[i].Area()
	}
	return func() Vec {
		rect := rects[RandIndex(areas)]
		return Vec{
			X: rand.Float64()*rect.W + rect.X,
			Y: rand.Float64()*rect.H + rect.Y,
		}
	}
}

// Returns a uniformaly distributed radius between minR and maxR.
func circleRadius(minR, maxR float64) float64 {
	if maxR == minR {
		return maxR
	}
	unitMin := minR / maxR
	unitMin *= unitMin
	return math.Sqrt(rand.Float64()*(1-unitMin)+unitMin) * maxR
}
