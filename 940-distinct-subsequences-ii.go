package main

func distinctSubseqII(s string) int {
	const mod = 1000000007
	endsWithCount := make([]int, 26)
	var result int
	for _, c := range s {
		i := int(c - 'a')
		prevCount := endsWithCount[i]
		endsWithCount[i] = (result + 1) % mod
		result += endsWithCount[i] - prevCount // 增加本次新增加的，减去重复的
		result = (result + mod) % mod
	}
	return result
}
