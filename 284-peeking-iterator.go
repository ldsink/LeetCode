package main

type PeekingIterator struct {
	hasCache bool
	item     int
	iter     *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasCache || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.hasCache {
		this.hasCache = false
		return this.item
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.hasCache {
		return this.item
	}
	this.hasCache = true
	this.item = this.iter.next()
	return this.item
}
