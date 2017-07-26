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

// RandNum returns a NumGen that returns a uniform random number between min and max.
func RandNum(min, max float64) NumGen {
	width := max - min
	return func() float64 {
		return rand.Float64()*width + min
	}
}

// RandRadius returns a NumGen that returns a uniform circle radius between minR and maxR.
func RandRadius(minR, maxR float64) NumGen {
	if maxR == 0 || maxR == minR {
		return func() float64 {
			return maxR
		}
	}
	unitMin := minR / maxR
	unitMin *= unitMin
	return func() float64 {
		return math.Sqrt(rand.Float64()*(1-unitMin)+unitMin) * maxR
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
//
func OffsetVec(gen VecGen, offset VecGen) VecGen {
	return func() Vec {
		return gen().Plus(offset())
	}
}

// RandVecCircle returns a VecGen that will generate a random vector within the given radii.
// A negative radius results in undefined behavior.
func RandVecCircle(minRadius, maxRadius float64) VecGen {
	return func() Vec {
		return RandVec().Times(circleRadius(minRadius, maxRadius))
	}
}

// RandVecArc returns a VecGen that will generate a random vector within the slice of a
// circle defined by the parameters. The radians are relative to the +x axis.
// A negative radius results in undefined behavior.
func RandVecArc(minRadius, maxRadius, minRadians, maxRadians float64) VecGen {
	if maxRadians < minRadians {
		minRadians, maxRadians = maxRadians, minRadians
	}
	return func() Vec {
		r := circleRadius(minRadius, maxRadius)
		rad := rand.Float64()*(maxRadians-minRadians) + minRadians
		return Vec{X: r}.Rotated(rad)
	}
}

// // RandVecRect returns a VecGen that will generate a random vector within the given Rect.
// func RandVecRect(rect Rect) VecGen {
// 	return func() Vec {
// 		return Vec{
// 			X: rand.Float64()*rect.W + rect.X,
// 			Y: rand.Float64()*rect.H + rect.Y,
// 		}
// 	}
// }
//
// // RandVecRects returns a VecGen that will generate a random vector that is uniformly
// // distributed between all the given rects. If the slice given is empty then the zero
// // vector is returned.
// func RandVecRects(rects []Rect) VecGen {
// 	if len(rects) == 0 {
// 		return func() Vec { return Vec{} }
// 	}
// 	areas := make([]float64, len(rects))
// 	for i := range rects {
// 		areas[i] = rects[i].Area()
// 	}
// 	return func() Vec {
// 		rect := rects[wrand.SelectIndex(areas)]
// 		return Vec{
// 			X: rand.Float64()*rect.W + rect.X,
// 			Y: rand.Float64()*rect.H + rect.Y,
// 		}
// 	}
// }

// Returns a uniformaly distributed radius between minR and maxR.
func circleRadius(minR, maxR float64) float64 {
	if maxR == 0 || maxR == minR {
		return maxR
	}
	unitMin := minR / maxR
	unitMin *= unitMin
	return math.Sqrt(rand.Float64()*(1-unitMin)+unitMin) * maxR
}
