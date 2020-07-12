package main

type NumArray struct {
	nums  []int
	sums  []int
	delta [][2]int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, j := len(nums)-1, 0; i >= 0; i-- {
		j += nums[i]
		sums[i] = j
	}
	var delta [][2]int
	return NumArray{nums: nums, sums: sums, delta: delta}
}

func (this *NumArray) Update(i int, val int) {
	delta := val - this.nums[i]
	this.nums[i] = val
	this.delta = append(this.delta, [2]int{i, delta})
}

func (this *NumArray) SumRange(i int, j int) int {
	result := this.sums[i] - this.sums[j+1]
	for _, d := range this.delta {
		if i <= d[0] && d[0] <= j {
			result += d[1]
		}
	}
	return result
}
