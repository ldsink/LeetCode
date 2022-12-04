package main

import (
	"math"
	"sort"
)

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	var toppings []int
	inSet := make(map[int]bool)
	var addTopping func(int, int)
	addTopping = func(idx, cost int) {
		if idx == len(toppingCosts) { // 到结尾了，添加这种配方方案
			if !inSet[cost] { // 去重，不重复添加
				inSet[cost] = true
				toppings = append(toppings, cost)
			}
			return
		}
		if cost < target {
			for i := 0; i < 3; i++ {
				addTopping(idx+1, cost+i*toppingCosts[idx])
			}
		} else { // 已经超过 target, 直接添加方案，减少无效枚举
			addTopping(len(toppingCosts), cost)
		}
	}
	addTopping(0, 0)
	sort.Ints(toppings) // 排序，方便后续二分查找配料方案

	findIndex := func(val int) int {
		start, end := 0, len(toppings)-1
		for middle := (start + end) >> 1; start != end; middle = (start + end) >> 1 {
			if val < toppings[middle+1] {
				end = middle
			} else {
				start = middle + 1
			}
		}
		return start
	}
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}

	result, diff := math.MaxInt, math.MaxInt
	for _, cost := range baseCosts {
		val := target - cost
		if val <= 0 { //冰激凌基料已经大于价格，不用再加配料
			if diff > -val {
				result = cost
				diff = -val
			}
			continue
		}
		idx := findIndex(val)
		if idx+1 < len(toppings) && diff > abs(val-toppings[idx+1]) {
			result = cost + toppings[idx+1]
			diff = abs(target - result)
		}
		if diff >= abs(val-toppings[idx]) {
			result = cost + toppings[idx]
			diff = abs(target - result)
		}
	}
	return result
}
