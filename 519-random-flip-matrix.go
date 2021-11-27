package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	m, n, max int
	used      map[int]bool
}

func Constructor(m int, n int) Solution {
	return Solution{m: m, n: n, max: m * n, used: make(map[int]bool)}
}

func (this *Solution) Flip() []int {
	var result []int
	rand.Seed(time.Now().UnixNano())
	for p := rand.Intn(this.max); ; {
		if this.used[p] {
			p = (p + 1) % this.max
			continue
		}
		this.used[p] = true
		i, j := p/this.n, p%this.n
		result = append(result, i)
		result = append(result, j)
		break
	}
	return result
}

func (this *Solution) Reset() {
	this.used = make(map[int]bool)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
