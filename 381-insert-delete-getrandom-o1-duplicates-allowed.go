package main

import (
	"crypto/rand"
	"math/big"
)

type CollectionNode struct {
	val, idx int
}

type RandomizedCollection struct {
	nodes []*CollectionNode
	set   map[int][]*CollectionNode
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{set: make(map[int][]*CollectionNode)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	_, ok := this.set[val]
	node := CollectionNode{val: val}
	this.set[val] = append(this.set[val], &node)
	this.nodes = append(this.nodes, &node)
	node.idx = len(this.nodes) - 1
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	nodes, ok := this.set[val]

	// val 不存在的情况
	if !ok {
		return false
	}

	// 删除节点
	last := nodes[len(nodes)-1]
	this.set[val] = this.set[val][:len(nodes)-1]
	if len(this.set[val]) == 0 {
		delete(this.set, val)
	}
	// 更新列表
	if last.idx != len(this.nodes)-1 {
		tail := this.nodes[len(this.nodes)-1]
		this.nodes[last.idx] = tail
		tail.idx = last.idx
	}
	this.nodes = this.nodes[:len(this.nodes)-1]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(this.nodes))))
	return this.nodes[int(n.Int64())].val

}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
