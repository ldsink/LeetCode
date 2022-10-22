package main

import (
	"math"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type job struct {
		end, profit int
	}
	jobMap := make(map[int][]job)
	for i := 0; i < len(startTime); i++ {
		jobMap[startTime[i]] = append(jobMap[startTime[i]], job{end: endTime[i], profit: profit[i]})
	}
	nodes := []int{math.MaxInt}
	for start := range jobMap {
		nodes = append(nodes, start)
	}
	sort.Ints(nodes)
	maxProfit := make([]int, len(nodes))
	for i := 0; i < len(maxProfit)-1; i++ {
		if i > 0 && maxProfit[i] < maxProfit[i-1] {
			maxProfit[i] = maxProfit[i-1]
		}
		for _, j := range jobMap[nodes[i]] {
			// 找到下一个节点位置
			next, end := i+1, len(nodes)-1
			for next != end {
				middle := (next + end) >> 1
				if j.end > nodes[middle] {
					next = middle + 1
				} else {
					end = middle
				}
			}
			// 更新下一个节点的最大收益
			if maxProfit[next] < maxProfit[i]+j.profit {
				maxProfit[next] = maxProfit[i] + j.profit
			}
		}
	}
	return maxProfit[len(nodes)-1]
}
