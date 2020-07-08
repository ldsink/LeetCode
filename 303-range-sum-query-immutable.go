package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	var sum int
	for ; i <= j; i++ {
		sum += this.nums[i]
	}
	return sum
}
