package main

type StockSpanner struct {
	prices, days []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.prices) == 0 || price < this.prices[len(this.prices)-1] {
		this.prices = append(this.prices, price)
		this.days = append(this.days, 1)
		return 1
	}
	idx := len(this.prices) - 1
	for ; idx >= 0 && this.prices[idx] <= price; idx -= this.days[idx] {
	}
	day := len(this.prices) - idx
	this.prices = append(this.prices, price)
	this.days = append(this.days, day)
	return day
}
