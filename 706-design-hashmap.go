package main

type MyHashMap struct {
	data map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{data: make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if val, ok := this.data[key]; ok {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.data, key)
}
