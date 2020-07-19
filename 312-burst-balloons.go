package main

func search(left, right int, balloons []int, coins [][]int) {
	if left+1 == right || coins[left][right] != 0 {
		return
	}
	for i := left + 1; i < right; i++ {
		current := balloons[left] * balloons[i] * balloons[right]
		search(left, i, balloons, coins)
		current += coins[left][i]
		search(i, right, balloons, coins)
		current += coins[i][right]
		if coins[left][right] < current {
			coins[left][right] = current
		}
	}
}

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	balloons := []int{1}
	balloons = append(balloons, nums...)
	balloons = append(balloons, 1)
	coins := make([][]int, len(balloons))
	for i := 0; i < len(balloons); i++ {
		coins[i] = make([]int, len(balloons))
	}
	search(0, len(balloons)-1, balloons, coins)
	return coins[0][len(balloons)-1]
}
