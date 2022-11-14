package main

import (
	"math"
	"sort"
)

func splitArraySameAverage(nums []int) bool {
	round := func(n float64) float64 {
		return math.Floor(n + 0.5)
	}
	abs := func(a float64) float64 {
		if a >= 0 {
			return a
		}
		return -a
	}

	sort.Ints(nums) // 排序，保持数组有序，相同的数值会在一起，方便剪枝
	var sum int
	for _, num := range nums {
		sum += num
	}
	avg := float64(sum) / float64(len(nums))

	cache := make(map[[3]int]bool)
	var find func(int, int, int) bool
	find = func(idx, count, val int) bool {
		if count == 0 {
			return val == 0
		} else if idx == len(nums) {
			return false
		}
		key := [3]int{idx, count, val}
		if v, ok := cache[key]; ok {
			return v
		}
		if idx+count == len(nums) {
			var sum int
			for i := idx; i < len(nums); i++ {
				sum += nums[i]
			}
			v := sum == val
			cache[key] = v
			return v
		}
		if val < nums[idx]*count { // 递增的数组，后面的值一定大于 val，那可以提前退出
			cache[key] = false
			return false
		}
		v := find(idx+1, count-1, val-nums[idx])
		if !v {
			v = find(idx+1, count, val)
		}
		cache[key] = v
		return v
	}

	// i 表示 i 个数字一组，可以凑成整数。这样可以尝试是否可以区分出 i 个数凑出一个合法的数值
	for i := 1; i <= len(nums)/2; i++ {
		val := avg * float64(i)
		num := int(round(val))
		if abs(float64(num)-val) < 1e-7 {
			if find(0, i, num) {
				return true
			}
		}
	}
	return false
}
