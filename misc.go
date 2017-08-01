package geo

import (
	"math"
	"math/rand"
	"sort"
)

// Clamp returns n if it is within [min(a, b), max(a, b)] otherwise it returns the closer
// of a and b.
func Clamp(n, a, b float64) float64 {
	if a > b {
		a, b = b, a
	}
	if n < a {
		return a
	}
	if n > b {
		return b
	}
	return n
}

// RandIndex takes a list of weights and returns an index with a probability corresponding
// to the relative weight of each index. Behavior is undefined if weights is the empty list.
// A weight of 0 will never be selected unless all are 0, in which case all indices have
// equal probability. Negative weights are treated as 0.
func RandIndex(weights []float64) int {
	cumWeights := make([]float64, len(weights))
	cumWeights[0] = weights[0]
	for i, w := range weights {
		if i > 0 {
			cumWeights[i] = cumWeights[i-1] + math.Max(w, 0)
		}
	}

	if cumWeights[len(weights)-1] == 0.0 {
		return rand.Intn(len(weights))
	}

	rnd := rand.Float64() * cumWeights[len(weights)-1]
	return sort.SearchFloat64s(cumWeights, rnd)
}

// Map takes a number n in the range [a1, b1] and remaps it be in the range [a2, b2], with
// its relative position within the range preserved. E.g if n is half way between a1 and b1
// then the value returned will be half way between a2 and b2. If a1 == b1 then +Inf is returned.
func Map(n, a1, b1, a2, b2 float64) float64 {
	range1 := b1 - a1
	range2 := b2 - a2
	offset := n - a1
	percent := offset / range1
	newOffset := percent * range2
	return a2 + newOffset
}

// Mod is the modulus operator. Unlike the Mod in the math package this wraps negative
// numbers around to the positive axis.
// e.g. if b is 3
//  a       -5 -4 -3 -2 -1 0 1 2 3 4 5
//  return   1  2  0  1  2 0 1 2 3 1 2
func Mod(a, b float64) float64 {
	return math.Mod(math.Mod(a, b)+b, b)
}

// Shake takes a current time t, from 0 to duration, the maximum amplitude and the frequency
// of the displacement, and a falloff function to control how the shaking dies off. The
// return Vec is the offset to use at time t. The seed value is used to vary the shake.
// But the same seed should be used for an entire shake cycle. It's purpose is to get different
// shake patterns when the other parameters are the same.
func Shake(seed, t, duration, amplitude, frequency float64, falloff EaseFn) Vec {
	t = Clamp(t, 0, duration) / duration
	amplitude *= 1 - falloff(t)
	return ShakeConst(seed, t, amplitude, frequency)
}

// ShakeConst produces a constant shake with no falloff or duration. It takes a maximum
// amplitude and the frequency of the displacement. The return Vec is the offset to use
// at time t. The seed value is used to vary the shake. But the same seed should be used for
// an entire shake cycle. It's purpose is to get different shake patterns when the other
// parameters are the same.
func ShakeConst(seed, t, amplitude, frequency float64) Vec {
	len := Map(Perlin(t*frequency, seed, seed), 0, 1, -1, 1) * amplitude
	angle := Map(Perlin(seed, t*frequency, seed), 0, 1, 0, 2*math.Pi)
	return VecLA(len, angle)
}
