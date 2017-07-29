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
