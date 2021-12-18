package main

func findJudge(n int, trust [][]int) int {
	count := make([]int, n+1)      // 被其他人的信任次数
	trustFlag := make([]bool, n+1) // 信任别人的标记
	for _, arr := range trust {
		a, b := arr[0], arr[1]
		count[b]++
		trustFlag[a] = true
	}
	for i := 1; i <= n; i++ {
		if count[i] == n-1 && trustFlag[i] == false {
			return i
		}
	}
	return -1
}
