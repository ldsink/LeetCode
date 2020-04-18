package main

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
	if node, ok := this.nodes[key]; !ok {
		return -1
	} else if node != this.tail {
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
	return this.tail.Value
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
