package main

import "strconv"

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	for start, length, count := 1, 1, 0; ; start, length = start*10, length+1 {
		end := start*10 - 1
		total := (end - start + 1) * length
		// length 位的数字全部加起来都没到 n，直接相加到下一位
		if count+total < n {
			count += total
			continue
		}
		num := start + ((n - count - 1) / length)
		s := strconv.Itoa(num)
		idx := (n - count - 1) % length
		num, _ = strconv.Atoi(s[idx : idx+1])
		return num
	}
}
