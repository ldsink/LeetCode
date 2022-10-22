package main

import (
	"math"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type job struct {
		start, end, profit, maxProfit int
	}
	var jobs []job
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, job{start: startTime[i], end: endTime[i], profit: profit[i]})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start || (jobs[i].start == jobs[j].start && jobs[i].end < jobs[j].end)
	})
	jobs = append(jobs, job{start: math.MaxInt})
	for i := 0; i < len(jobs)-1; i++ {
		if i > 0 && jobs[i].maxProfit < jobs[i-1].maxProfit {
			jobs[i].maxProfit = jobs[i-1].maxProfit
		}
		// 找到下一个节点位置
		next, end := i+1, len(jobs)-1
		for next != end {
			middle := (next + end) >> 1
			if jobs[i].end > jobs[middle].start {
				next = middle + 1
			} else {
				end = middle
			}
		}
		// 更新下一个节点的最大收益
		if jobs[next].maxProfit < jobs[i].maxProfit+jobs[i].profit {
			jobs[next].maxProfit = jobs[i].maxProfit + jobs[i].profit
		}
	}
	return jobs[len(jobs)-1].maxProfit
}
