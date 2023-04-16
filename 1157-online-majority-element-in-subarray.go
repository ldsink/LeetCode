package main

import "math/rand"

type MajorityChecker struct {
	indexes map[int][]int
	arr     []int
}

func Constructor(arr []int) MajorityChecker {
	indexes := make(map[int][]int)
	for i, v := range arr {
		indexes[v] = append(indexes[v], i)
	}
	return MajorityChecker{indexes: indexes, arr: arr}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	if right-left <= 20 { // 长度小于 20，直接搜索
		count := make(map[int]int)
		for ; left <= right; left++ {
			count[this.arr[left]]++
			if count[this.arr[left]] >= threshold {
				return this.arr[left]
			}
		}
		return -1
	}
	for t, searched := 20, make(map[int]bool); t >= 0; t-- { // 否则随机取出 20 个数，验证是否存在众数
		if v := this.arr[left+rand.Intn(right-left)]; !searched[v] {
			searched[v] = true
			arr := this.indexes[v]
			if len(arr) < threshold || arr[0] > right || arr[len(arr)-1] < left {
				continue
			}
			start, end := 0, len(arr)-1
			for start != end {
				middle := (start + end) >> 1
				if arr[middle] < left {
					start = middle + 1
				} else {
					end = middle
				}
			}
			if arr[start] < left || start+threshold-1 >= len(arr) || arr[start+threshold-1] > right {
				continue
			}
			return v
		}
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
