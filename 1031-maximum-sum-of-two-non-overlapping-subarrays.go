package main

import (
	"sort"
)

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	type Status struct {
		idx, sum int
	}
	getSumList := func(l int) (r []Status) {
		var sum int
		for i := 0; i < l; i++ {
			sum += nums[i]
		}
		r = append(r, Status{idx: 0, sum: sum})
		for i := l; i < len(nums); i++ {
			sum += nums[i] - nums[i-l]
			r = append(r, Status{idx: i - l + 1, sum: sum})
		}
		sort.Slice(r, func(i, j int) bool {
			return r[i].sum > r[j].sum || (r[i].sum == r[j].sum && r[i].idx < r[j].idx)
		})
		return
	}
	firstSum, secondSum := getSumList(firstLen), getSumList(secondLen)
	var result int
	for _, item := range firstSum {
		for _, next := range secondSum {
			if result >= item.sum+next.sum {
				break
			}
			a, b, c, d := item.idx, item.idx+firstLen-1, next.idx, next.idx+secondLen-1
			if a > c { // 确保 a 代表最左边的点
				a, b, c, d = c, d, a, b
			}
			if b < c && result < item.sum+next.sum {
				result = item.sum + next.sum
				break
			}
		}
	}
	return result
}
