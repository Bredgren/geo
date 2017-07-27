package geo

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
