package main

func preimageSizeFZF(k int) int {
	trailingZeroes := func(n int) (result int) {
		for i := 5; n/i != 0; i *= 5 {
			result += n / i
		}
		return result
	}
	const maxNUm = 1 << 32 // max 1073741816 zeros, > 10^9
	start, end, anyIndex := 0, maxNUm, -1
	for start < end {
		middle := (start + end) >> 1
		zeros := trailingZeroes(middle)
		if zeros == k {
			anyIndex = middle
			break
		} else if zeros > k {
			end = middle
		} else {
			start = middle + 1
		}
	}
	if anyIndex == -1 { // 不存在数满足 k 个零结尾
		return 0
	}
	// 二分法找到最左边的和最右边的
	for start, end = 0, anyIndex; start < end; {
		middle := (start + end) >> 1
		zeros := trailingZeroes(middle)
		if k <= zeros {
			end = middle
		} else {
			start = middle + 1
		}
	}
	left := start
	for start, end = anyIndex, maxNUm; start < end; {
		middle := (start+end)>>1 + 1
		zeros := trailingZeroes(middle)
		if k >= zeros {
			start = middle
		} else {
			end = middle - 1
		}
	}
	right := start
	return right - left + 1
}
