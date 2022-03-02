package main

import (
	"math"
	"strconv"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func getMinDiff(diff, num int, n, result, curr string) (int, string) {
	k, _ := strconv.Atoi(curr)
	if d := abs(k - num); n != curr && (diff > d || (diff == d && k < num)) {
		diff, result = d, curr
	}
	return diff, result
}

func nearestPalindromic(n string) string {
	num, _ := strconv.Atoi(n)
	diff := math.MaxInt64
	var result string

	//将前一半填入后一半
	curr := []rune(n)
	middle := len(n) >> 1
	if middle > 0 {
		for i := 0; i <= middle; i++ {
			curr[len(n)-1-i] = curr[i]
		}
		diff, result = getMinDiff(diff, num, n, result, string(curr))
	}

	// 中间 +1
	middleRune := curr[middle]
	if middleRune < '9' {
		if len(n)%2 == 1 { // 奇数长度的处理情况
			curr[middle] = middleRune + 1
		} else {
			curr[middle] = middleRune + 1
			curr[middle-1] = middleRune + 1
		}
		diff, result = getMinDiff(diff, num, n, result, string(curr))
	}

	// 中间 -1
	if middleRune > '0' {
		if len(n)%2 == 1 { // 奇数长度的处理情况
			curr[middle] = middleRune - 1
		} else {
			curr[middle] = middleRune - 1
			curr[middle-1] = middleRune - 1
		}
		diff, result = getMinDiff(diff, num, n, result, string(curr))
	}

	//下边界
	if len(n) > 1 {
		curr = []rune{}
		for i := 0; i < len(n)-1; i++ {
			curr = append(curr, '9')
		}
		diff, result = getMinDiff(diff, num, n, result, string(curr))
	}

	//上边界
	curr = []rune{'1'}
	for i := 0; i < len(n)-1; i++ {
		curr = append(curr, '0')
	}
	curr = append(curr, '1')
	diff, result = getMinDiff(diff, num, n, result, string(curr))

	return result
}
