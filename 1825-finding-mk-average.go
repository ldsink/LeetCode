package main

import "github.com/emirpasic/gods/trees/redblacktree"

type MKAverage struct {
	small, middle, large          *redblacktree.Tree
	queue                         []int
	m, k, sum, smallLen, largeLen int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{small: redblacktree.NewWithIntComparator(), middle: redblacktree.NewWithIntComparator(), large: redblacktree.NewWithIntComparator(), m: m, k: k}
}

func (this *MKAverage) AddElement(num int) {
	merge := func(rbt *redblacktree.Tree, key, value int) {
		if v, ok := rbt.Get(key); ok {
			next := v.(int) + value
			if next == 0 {
				rbt.Remove(key)
			} else {
				rbt.Put(key, next)
			}
		} else {
			rbt.Put(key, value)
		}
	}

	if this.small.Empty() || num <= this.small.Right().Key.(int) {
		merge(this.small, num, 1)
		this.smallLen++
	} else if this.large.Empty() || num >= this.large.Left().Key.(int) {
		merge(this.large, num, 1)
		this.largeLen++
	} else {
		merge(this.middle, num, 1)
		this.sum += num
	}
	this.queue = append(this.queue, num)
	if len(this.queue) > this.m {
		x := this.queue[0]
		this.queue = this.queue[1:]
		if _, ok := this.small.Get(x); ok {
			merge(this.small, x, -1)
			this.smallLen--
		} else if _, ok := this.large.Get(x); ok {
			merge(this.large, x, -1)
			this.largeLen--
		} else {
			merge(this.middle, x, -1)
			this.sum -= x
		}
	}
	for ; this.smallLen > this.k; this.smallLen-- {
		x := this.small.Right().Key.(int)
		merge(this.small, x, -1)
		merge(this.middle, x, 1)
		this.sum += x
	}
	for ; this.largeLen > this.k; this.largeLen-- {
		x := this.large.Left().Key.(int)
		merge(this.large, x, -1)
		merge(this.middle, x, 1)
		this.sum += x
	}
	for ; this.smallLen < this.k && !this.middle.Empty(); this.smallLen++ {
		x := this.middle.Left().Key.(int)
		merge(this.middle, x, -1)
		this.sum -= x
		merge(this.small, x, 1)
	}
	for ; this.largeLen < this.k && !this.middle.Empty(); this.largeLen++ {
		x := this.middle.Right().Key.(int)
		merge(this.middle, x, -1)
		this.sum -= x
		merge(this.large, x, 1)
	}

}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.queue) < this.m {
		return -1
	}
	return this.sum / (this.m - 2*this.k)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
