package geo

import (
	"math"
	"testing"
)

func TestClamp(t *testing.T) {
	cases := []struct {
		n, a, b, want float64
	}{
		{0, 0, 0, 0},
		{0, 1, 2, 1},
		{0, 2, 1, 1},
		{3, 1, 2, 2},
		{3, 2, 1, 2},
		{1.2, 1, 2, 1.2},
		{1.2, 2, 1, 1.2},
	}
	for i, c := range cases {
		got := Clamp(c.n, c.a, c.b)
		if got != c.want {
			t.Errorf("case %d: got: %f, want: %f", i, got, c.want)
		}
	}
}

func TestMap(t *testing.T) {
	cases := []struct {
		n, a1, b1, a2, b2, want float64
	}{
		{0, 0, 1, 5, 10, 5},
		{1, 0, 1, 5, 10, 10},
		{0.5, 0, 1, 5, 11, 8},
		{1, 0, 0, 5, 11, math.Inf(1)},
		{1, 0, 2, 5, 5, 5},
	}
	for i, c := range cases {
		got := Map(c.n, c.a1, c.b1, c.a2, c.b2)
		if got != c.want {
			t.Errorf("case %d: got: %f, want: %f", i, got, c.want)
		}
	}
}

func TestMod(t *testing.T) {
	cases := []struct {
		a, b, want float64
	}{
		{-4, 3, 2},
		{-3, 3, 0},
		{-2, 3, 1},
		{-1, 3, 2},
		{0, 3, 0},
		{1, 3, 1},
		{2, 3, 2},
		{3, 3, 0},
		{4, 3, 1},
		{1, 0, 0},
	}
	for i, c := range cases {
		got := Mod(c.a, c.b)
		if got != c.want {
			t.Errorf("case %d: got: %f, want: %f", i, got, c.want)
		}
	}
}

func TestRandIndex(t *testing.T) {
	cases := [][]float64{
		{0.0},
		{0.0, 0.0},
		{0.0, 0.0, 1.0},
		{1.0},
		{1.0, 2.0},
		{1.0, 4.0},
		{1.0, 4.0, 2.0},
		{1.5, 2.5},
		{-1.0, -1.0},
	}

	for _, c := range cases {
		checkSelectIndexFreq(c)
	}
}

func checkSelectIndexFreq(list []float64) {
	counts := make(map[int]int)
	count := 1000
	for i := 0; i < count; i++ {
		index := RandIndex(list)
		counts[index]++
	}

	// fmt.Println("Check", list)
	// for v, c := range counts {
	// 	fmt.Printf(" %d: %f\n", v, float64(c)/float64(count))
	// }
}

func TestI2E2(t *testing.T) {
	i1, i2 := 1, 2
	f1, f2 := I2F2(i1, i2)
	if f1 != 1.0 || f2 != 2.0 {
		t.Errorf("got %f, %f, want %f, %f", f1, f2, 1.0, 2.0)
	}
}
