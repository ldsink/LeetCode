package main

import "strings"

type MapSum struct {
	data map[string]int
}

func Constructor() MapSum {
	return MapSum{data: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.data[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	var result int
	for k, v := range this.data {
		if strings.HasPrefix(k, prefix) {
			result += v
		}
	}
	return result
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
