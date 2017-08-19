package geo

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	e = 1e-10
)

func init() {
	rand.Seed(time.Now().Unix())
}

func check(t *testing.T, name string, fn func() float64, want float64) {
	got := fn()
	if got != want {
		t.Errorf("%s: got %f, want %f", name, got, want)
	}
}

func check2(t *testing.T, name string, fn func() (float64, float64), want1, want2 float64) {
	got1, got2 := fn()
	if got1 != want1 || got2 != want2 {
		t.Errorf("%s: got %f, %f, want %f, %f", name, got1, got2, want1, want2)
	}
}

func check3(t *testing.T, name string, fn func() (float64, float64, float64), want1, want2, want3 float64) {
	got1, got2, got3 := fn()
	if got1 != want1 || got2 != want2 || got3 != want3 {
		t.Errorf("%s: got %f, %f, %f, want %f, %f, %f", name, got1, got2, got3, want1, want2, want3)
	}
}

func checkVec(t *testing.T, name string, fn func() Vec, want Vec) {
	got := fn()
	if got != want {
		t.Errorf("%s: got %s, want %s", name, got, want)
	}
}

func intListEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}

func fEqual(a, b float64) bool {
	return (math.IsInf(a, 1) && math.IsInf(b, 1)) ||
		(math.IsInf(a, -1) && math.IsInf(b, -1)) ||
		(math.IsNaN(a) && math.IsNaN(b)) ||
		math.Abs(a-b) < e
}
