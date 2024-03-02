package main

func maxProfit(prices []int, fee int) int {
	bestResult := make([]int, len(prices)+1)
	processed := make([]bool, len(prices)+1)
	processed[len(processed)-1] = true
	var findBestResult func(int)
	findBestResult = func(i int) {
		if processed[i] {
			return
		}
		processed[i] = true
		if i == len(prices)-1 { // 到最后一个了，一定是 0，没有收益
			return
		} else if i == len(prices)-2 {
			if p := prices[len(prices)-1] - prices[len(prices)-2] - fee; p > 0 {
				bestResult[i] = p
			}
			return
		}
		// 买入股票的情况
		for j, k := i+1, prices[i]+fee; j < len(prices); j++ {
			if prices[j] <= prices[i] { // 后面的价格比我底，最佳买入点一定在后面
				findBestResult(j)
				if bestResult[i] < bestResult[j] {
					bestResult[i] = bestResult[j]
				}
				break
			}
			if prices[j] > k {
				findBestResult(j + 1)
				if c := prices[j] - k + bestResult[j+1]; bestResult[i] < c {
					bestResult[i] = c
				}
			}
		}
	}
	findBestResult(0)
	return bestResult[0]
}
