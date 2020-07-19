package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var buy, buyPrices, sell, sellPrices, cooldown [2]int
	buyPrices[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		curr := i % 2
		prev := curr ^ 1
		// buy
		buy[curr] = buy[prev]
		buyPrices[curr] = buyPrices[prev]
		if buyPrices[curr] > prices[i] {
			buyPrices[curr] = prices[i]
		}
		if cooldown[prev]-prices[i] >= buy[curr]-buyPrices[curr] {
			buy[curr] = cooldown[prev]
			buyPrices[curr] = prices[i]
		}
		// sell
		sell[curr] = sell[prev]
		sellPrices[curr] = sellPrices[prev]
		if delta := buy[prev] + prices[i] - buyPrices[prev]; sell[curr] < delta {
			sell[curr] = delta
		}
		// cooldown
		cooldown[curr] = cooldown[prev]
		if cooldown[curr] < sell[prev] {
			cooldown[curr] = sell[prev]
		}
	}
	return sell[(len(prices)-1)%2]
}
