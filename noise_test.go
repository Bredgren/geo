package geo

import (
	"math"
	"math/rand"
	"testing"
)

func TestPerlin(t *testing.T) {
	cases := []struct {
		x, y, z, want float64
	}{
		// This fist case from https://rosettacode.org/wiki/Perlin_noise#Go
		{3.14, 42, 7, Map(0.13691995878400012, -1, 1, 0, 1)},
	}

	for i, c := range cases {
		got := Perlin(c.x, c.y, c.z)
		if got != c.want {
			t.Errorf("case %d: got %f, want %f", i, got, c.want)
		}
	}

	for i := 0; i < 100; i++ {
		x, y, z := rand.Float64()*1000, rand.Float64()*1000, rand.Float64()*1000
		got := Perlin(x, y, z)
		if got < 0 || got > 1 {
			t.Errorf("out of range: got %f, x %f y %f, z %f", got, x, y, z)
		}
	}
}

func TestPerlinOctave(t *testing.T) {
	a := Perlin(1.2, 3.4, 4.5)
	b := Perlin(2.4, 6.8, 9.0)

	got := PerlinOctave(1.2, 3.4, 4.5, 1, 1)
	if got != a {
		t.Errorf("got %f, want %f", got, a)
	}

	got = PerlinOctave(1.2, 3.4, 4.5, 2, 1)
	if math.Abs(got-(a+b)) < e {
		t.Errorf("got %f, want %f", got, a+b)
	}

	got = PerlinOctave(1.2, 3.4, 4.5, 2, 0.5)
	if math.Abs(got-(a+b)/2) < e {
		t.Errorf("got %f, want %f", got, (a+b)/2)
	}
}
