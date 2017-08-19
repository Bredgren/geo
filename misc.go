package geo

import (
	"math"
	"math/rand"
	"sort"
	"time"
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
// numbers around to the positive axis. If b is 0 then this function returns 0.
// e.g. if b is 3
//  a       -5 -4 -3 -2 -1 0 1 2 3 4 5
//  return   1  2  0  1  2 0 1 2 0 1 2
func Mod(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return math.Mod(math.Mod(a, b)+b, b)
}

// I2F2 takes 2 ints and converts them to float64. e.g.
//  rect.SetTopLeft(I2F2(functionThatReturns2Ints()))
func I2F2(i1, i2 int) (f1, f2 float64) {
	return float64(i1), float64(i2)
}

// Shaker wraps the arguments to the shake functions for convenience and reusability.
type Shaker struct {
	// Seed can be used to change up the shaking behaviour, because all of the shake functions
	// are deterministic and often one wants it look different while keeping the other parameters
	// the same. Though if different values for t are always being used then changing the Seed
	// may not be necessary.
	Seed float64
	// StartTime is the time that the shaking starts. Reactivating a Shaker that has ended
	// should usually be as simple as updating the StartTime.
	StartTime time.Time
	// Duration is how long the shaking takes place.
	Duration time.Duration
	// Amplitude is the maximum length of the offset.
	Amplitude float64
	// Frequency controls how quickly the shaking happens.
	Frequency float64
	// Falloff modifies the amplitude over time, using StartTime and Duration. This makes
	// it easy to fade out the shaking.
	Falloff EaseFn
}

// EndTime returns the time that the shaking would end, if using the non-Const Shake functions.
func (s *Shaker) EndTime() time.Time {
	return s.StartTime.Add(s.Duration)
}

// Shake takes a current time t and returns an offset. The time t will be clamped between
// s.StartTime and s.StartTime + s.Duration. This function makes use of StartTime, Duration,
// and Falloff to change the amplitude of the offset over time. The length of the Vec returned
// is never greater than the amplitude.
func (s *Shaker) Shake(t time.Time) Vec {
	return s.shakeConst(t, s.amp(t))
}

// ShakeConst takes a current time t and returns an offset. The max amplitude of the offset
// is not varied over time so StartTime, Duration, and Falloff are not used.
func (s *Shaker) ShakeConst(t time.Time) Vec {
	return s.shakeConst(t, s.Amplitude)
}

// Shake1 is the same as Shake function but works in 1 dimension.
func (s *Shaker) Shake1(t time.Time) float64 {
	return s.shakeConst1(t, s.amp(t))
}

// ShakeConst1 is the same as the ShakeConst but works in 1 dimension.
func (s *Shaker) ShakeConst1(t time.Time) float64 {
	return s.shakeConst1(t, s.Amplitude)
}

func (s *Shaker) amp(t time.Time) float64 {
	dt := Clamp(t.Sub(s.StartTime).Seconds(), 0, s.Duration.Seconds()) / s.Duration.Seconds()
	return s.Amplitude * (1 - s.Falloff(dt))
}

func (s *Shaker) shakeConst(t time.Time, amplitude float64) Vec {
	dt := time.Duration(t.UnixNano()).Seconds()
	len := Map(Perlin(dt*s.Frequency, s.Seed, s.Seed), 0, 1, -1, 1) * amplitude
	angle := Map(Perlin(s.Seed, dt*s.Frequency, s.Seed), 0, 1, -math.Pi, math.Pi)
	return VecLA(len, angle)
}

func (s *Shaker) shakeConst1(t time.Time, amplitude float64) float64 {
	dt := time.Duration(t.UnixNano()).Seconds()
	return Map(Perlin(dt*s.Frequency, s.Seed, s.Seed), 0, 1, -1, 1) * amplitude
}
