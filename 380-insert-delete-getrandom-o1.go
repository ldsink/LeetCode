package main

import (
	"crypto/rand"
	"math/big"
)

type RandomizedSet struct {
	idx []int
	set map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.set[val] = len(this.idx)
	this.idx = append(this.idx, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.set[val]
	if !ok {
		return false
	}
	last := this.idx[len(this.idx)-1]
	this.idx[idx] = last
	this.idx = this.idx[:len(this.idx)-1]
	this.set[last] = idx
	delete(this.set, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(this.idx))))
	return this.idx[int(n.Int64())]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
