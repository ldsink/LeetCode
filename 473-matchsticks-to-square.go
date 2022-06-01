package main

import (
	"sort"
)

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	var edge int
	for _, l := range matchsticks {
		edge += l
	}
	if edge%4 != 0 {
		return false
	}
	edge /= 4

	// 动规+状态压缩。前 32 位代表当前的火柴索引，后 32 位记录每个火柴的放置情况。每个火柴 4 种情况 2bit 位，不超过 15 个火柴，最多 30bit 位。
	sort.Slice(matchsticks, func(i, j int) bool { return matchsticks[i] > matchsticks[j] })
	const offset = 32
	const fullBit = 0xffffffff
	var dp func(int, int, int, int, int) bool
	dp = func(key, a, b, c, d int) bool {
		idx := key >> offset
		if idx == len(matchsticks) {
			success := false
			if edge == a && a == b && b == c && c == d {
				success = true
			}
			return success
		}
		next := (idx + 1) << offset
		success := a+matchsticks[idx] <= edge && dp(next|(key&fullBit), a+matchsticks[idx], b, c, d)
		success = success || (b+matchsticks[idx] <= edge && dp(next|(1<<(idx*2)|(key&fullBit)), a, b+matchsticks[idx], c, d))
		success = success || (c+matchsticks[idx] <= edge && dp(next|(2<<(idx*2)|(key&fullBit)), a, b, c+matchsticks[idx], d))
		success = success || (d+matchsticks[idx] <= edge && dp(next|(3<<(idx*2)|(key&fullBit)), a, b, c, d+matchsticks[idx]))
		return success
	}
	return dp(0, 0, 0, 0, 0)
}
