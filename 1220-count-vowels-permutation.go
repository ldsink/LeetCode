package main

func countVowelPermutation(n int) int {
	const mod = 1000000007
	var count [2][5]int // 记录数量的数字，第二层 0-4 分别表示 a e i o u

	// 初始化 1 的情况
	for i := 0; i < 5; i++ {
		count[1][i] = 1
	}
	// 从 2 到 n 的情况
	for i := 2; i <= n; i++ {
		curr, prev := i&1, (i+1)&1
		// 每个元音 'a' 后面都只能跟着 'e'
		count[curr][0] = count[prev][1]
		// 每个元音 'e' 后面只能跟着 'a' 或者是 'i'
		count[curr][1] = (count[prev][0] + count[prev][2]) % mod
		// 每个元音 'i' 后面 不能 再跟着另一个 'i'
		count[curr][2] = (count[prev][0] + count[prev][1] + count[prev][3] + count[prev][4]) % mod
		// 每个元音 'o' 后面只能跟着 'i' 或者是 'u'
		count[curr][3] = (count[prev][2] + count[prev][4]) % mod
		// 每个元音 'u' 后面只能跟着 'a'
		count[curr][4] = count[prev][0]
	}
	var result int
	for i := 0; i < 5; i++ {
		result += count[n&1][i]
		result %= mod
	}
	return result
}
