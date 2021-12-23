package main

func longestDupSubstring(s string) string {
	a := 'a'
	var indexes [26][]int
	for i, c := range s {
		indexes[c-a] = append(indexes[c-a], i)
	}
	var idx, max int
	for i := 0; i < len(indexes); i++ {
		if len(indexes[i]) > 1 {
			idx, max = indexes[i][0], 1
			break
		}
	}
	if max == 0 {
		return ""
	}
	// 下面为至少有长度为 1 的子串的情况
	failed := make([]bool, len(s))                          // 某个位置明确不支持的长度
	for i, success := 2, true; success && i < len(s); i++ { // 当前可以成功的长度
		success = false
		beginIndex := make([]int, 26)
		for j := 0; j < len(s)-i; j++ {
			c := s[j] - uint8(a)
			beginIndex[c]++
			if failed[j] {
				continue // 如果 j 这个位置在更短的 i 的位置下失败过，那当前更加长的 i 也不可能
			}
			for k := beginIndex[c]; k < len(indexes[c]); k++ {
				if indexes[c][k]+i-1 >= len(s) {
					break
				}
				if s[j:j+i] == s[indexes[c][k]:indexes[c][k]+i] {
					success = true
					break
				}
			}
			if success {
				max, idx = i, j
				break
			}
			failed[j] = true // 在 i 这个长度下，这个位置是挂掉的
		}
	}
	return s[idx : idx+max]
}
