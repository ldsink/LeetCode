package main

import (
	"math/rand"
)

type Solution struct {
	o []int
	s []int
}

func Constructor(nums []int) Solution {
	shuffle := make([]int, len(nums))
	copy(shuffle, nums)
	return Solution{o: nums, s: shuffle}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.o
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.s), func(i, j int) { this.s[i], this.s[j] = this.s[j], this.s[i] })
	return this.s
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
