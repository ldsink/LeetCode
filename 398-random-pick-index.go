package main

import "math/rand"

type Solution struct {
	indexes map[int][]int
}

func Constructor(nums []int) Solution {
	indexes := make(map[int][]int)
	for i, n := range nums {
		indexes[n] = append(indexes[n], i)
	}
	return Solution{indexes: indexes}
}

func (this *Solution) Pick(target int) int {
	idx := rand.Intn(len(this.indexes[target]))
	return this.indexes[target][idx]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
