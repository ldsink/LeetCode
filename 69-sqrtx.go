package main

import "math"

func mySqrt(x int) int {
	r := math.Sqrt(float64(x))
	return int(r)
}
