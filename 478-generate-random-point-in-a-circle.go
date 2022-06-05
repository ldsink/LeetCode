package main

import (
	"math/rand"
)

type Solution struct {
	x, y, r float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{x: x_center, y: y_center, r: radius}
}

func (this *Solution) RandPoint() []float64 {
	for {
		if x, y := rand.Float64()*2-1, rand.Float64()*2-1; x*x+y*y < 1 {
			return []float64{this.x + x*this.r, this.y + y*this.r}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
