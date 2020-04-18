package main

import (
	"fmt"
)

type keyNode struct {
	Key   int
	Value int
	Prev  *keyNode
	Next  *keyNode
}

type LRUCache struct {
	capacity int
	head     *keyNode
	tail     *keyNode
	nodes    map[int]*keyNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		nodes:    make(map[int]*keyNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodes[key]
	if !ok {
		return -1
	}
	if node != this.tail {
		if node == this.head {
			this.head = this.head.Next
			this.head.Prev = nil
		} else {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
		}
		this.tail.Next = node
		node.Prev = this.tail
		this.tail = node
		this.tail.Next = nil
	}
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	// cache is empty
	if this.head == nil {
		node := keyNode{Key: key, Value: value}
		this.head = &node
		this.tail = &node
		this.nodes[key] = &node
		return
	}
	// key in cache
	if -1 != this.Get(key) {
		this.nodes[key].Value = value
		return
	}
	// add key to cache
	node := keyNode{Key: key, Value: value, Prev: this.tail}
	this.tail.Next = &node
	this.tail = &node
	this.nodes[key] = &node
	for ; len(this.nodes) > this.capacity; {
		delete(this.nodes, this.head.Key)
		this.head = this.head.Next
		this.head.Prev = nil
	}
}

//["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
//[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]

func main() {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	for i := 4; i > 0; i-- {
		fmt.Println(c.Get(i))
	}
	c.Put(5, 5)
	for i := 1; i < 6; i++ {
		fmt.Println(c.Get(i))
	}
}
