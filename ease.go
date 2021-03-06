package geo

import "math"

// Lerp linear interpolates from a to b by percent t, where t=0 returns a and
// t=1 returns b.
func Lerp(a, b, t float64) float64 {
	return (1-t)*a + t*b
}

// LerpVec linear interpolates from a to b by percent t, where t=0 returns a
// and t=1 returns b.
func LerpVec(a, b Vec, t float64) Vec {
	return Vec{X: Lerp(a.X, b.X, t), Y: Lerp(a.Y, b.Y, t)}
}

// Ease interpolates from a to b by percent t following the path defined by easefn.
func Ease(a, b, t float64, easefn EaseFn) float64 {
	return Lerp(a, b, easefn(t))
}

// EaseVec interpolates from a to b by percent t following the path defined by easefn.
func EaseVec(a, b Vec, t float64, easefn EaseFn) Vec {
	return Vec{X: Ease(a.X, b.X, t, easefn), Y: Ease(a.Y, b.Y, t, easefn)}
}

// EaseFn is a function that typically takes a time t between 0 and 1 and maps it to
// a different value between 0 and 1.
//
// Ease functions taken from https://gist.github.com/gre/1650294 and
// https://github.com/warrenm/AHEasing/blob/master/AHEasing/easing.c
type EaseFn func(t float64) float64

// EaseIn creates an EaseFn with the given power that eases into the destination.
func EaseIn(power float64) EaseFn {
	return func(t float64) float64 {
		return math.Pow(t, power)
	}
}

// EaseOut creates an EaseFn with the given power that eases out of the start.
func EaseOut(power float64) EaseFn {
	return func(t float64) float64 {
		return 1 - math.Abs(math.Pow(t-1, power))
	}
}

// EaseInOut creates an EaseFn with the given power that eases both at the start and the
// end.
func EaseInOut(power float64) EaseFn {
	return func(t float64) float64 {
		if t < 0.5 {
			return EaseIn(power)(t*2) / 2
		}
		return EaseOut(power)(t*2-1)/2 + 0.5
	}
}

var (
	EaseLinear = EaseIn(1)

	EaseInQuad    = EaseIn(2)
	EaseOutQuad   = EaseOut(2)
	EaseInOutQuad = EaseInOut(2)

	EaseInCubic    = EaseIn(3)
	EaseOutCubic   = EaseOut(3)
	EaseInOutCubic = EaseInOut(3)

	EaseInQuart    = EaseIn(4)
	EaseOutQuart   = EaseOut(4)
	EaseInOutQuart = EaseInOut(4)

	EaseInQuint    = EaseIn(5)
	EaseOutQuint   = EaseOut(5)
	EaseInOutQuint = EaseInOut(5)
)

func EaseInSine(t float64) float64 {
	return math.Sin((t-1)*math.Pi/2) + 1
}

func EaseOutSine(t float64) float64 {
	return math.Sin(t * math.Pi / 2)
}

func EaseInOutSine(t float64) float64 {
	return 0.5 * (1 - math.Cos(t*math.Pi))
}

func EaseInCirc(t float64) float64 {
	return 1 - math.Sqrt(1-(t*t))
}

func EaseOutCirc(t float64) float64 {
	return math.Sqrt((2 - t) * t)
}

func EaseInOutCirc(t float64) float64 {
	if t < 0.5 {
		return 0.5 * (1 - math.Sqrt(1-4*(t*t)))
	}
	return 0.5 * (math.Sqrt(-((2*t)-3)*((2*t)-1)) + 1)
}

func EaseInExpo(t float64) float64 {
	if t == 0 {
		return t
	}
	return math.Pow(2, 10*(t-1))
}

func EaseOutExpo(t float64) float64 {
	if t == 1 {
		return t
	}
	return 1 - math.Pow(2, -10*t)
}

func EaseInOutExpo(t float64) float64 {
	if t == 0 || t == 1 {
		return t
	}
	if t < 0.5 {
		return 0.5 * math.Pow(2, (20*t)-10)
	}
	return -0.5*math.Pow(2, (-20*t)+10) + 1
}

func EaseInElastic(t float64) float64 {
	return math.Sin(13*math.Pi/2*t) * math.Pow(2, 10*(t-1))
}

func EaseOutElastic(t float64) float64 {
	return math.Sin(-13*math.Pi/2*(t+1))*math.Pow(2, -10*t) + 1
}

func EaseInOutElastic(t float64) float64 {
	if t < 0.5 {
		return 0.5 * math.Sin(13*math.Pi/2*(2*t)) * math.Pow(2, 10*((2*t)-1))
	}
	return 0.5 * (math.Sin(-13*math.Pi/2*((2*t-1)+1))*math.Pow(2, -10*(2*t-1)) + 2)
}

func EaseInBack(t float64) float64 {
	return t*t*t - t*math.Sin(t*math.Pi)
}

func EaseOutBack(t float64) float64 {
	return 1 - EaseInBack(1-t)
}

func EaseInOutBack(t float64) float64 {
	if t < 0.5 {
		return 0.5 * EaseInBack(2*t)
	}
	return 0.5*EaseOutBack(2*t-1) + 0.5
}

func EaseInBounce(t float64) float64 {
	return 1 - EaseOutBounce(1-t)
}

func EaseOutBounce(t float64) float64 {
	if t < 4/11.0 {
		return (121 * t * t) / 16.0
	}
	if t < 8/11.0 {
		return (363 / 40.0 * t * t) - (99 / 10.0 * t) + 17/5.0
	}
	if t < 9/10.0 {
		return (4356 / 361.0 * t * t) - (35442 / 1805.0 * t) + 16061/1805.0
	}
	return (54 / 5.0 * t * t) - (513 / 25.0 * t) + 268/25.0
}

func EaseInOutBounce(t float64) float64 {
	if t < 0.5 {
		return 0.5 * EaseInBounce(t*2)
	}
	return 0.5*EaseOutBounce(t*2-1) + 0.5
}
