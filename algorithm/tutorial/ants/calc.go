package ants

import "math"

func Calc(l, n int, pos []int) (min int, max int) {
	for _, p := range pos {
		min = int(math.Max(float64(min), math.Min(float64(p), float64(l-p))))
		max = int(math.Max(float64(max), math.Max(float64(p), float64(l-p))))
	}
	return
}
