package main

import "fmt"

const failed = -1

func findIndex(number int, nums []int, start, end int) int {
	if start == end {
		if nums[start] == number {
			return start
		}
		return failed
	}
	middle := (start + end) >> 1
	if nums[middle] == number {
		return middle
	} else if nums[middle] > number {
		end = middle - 1
		if end < start {
			return failed
		}
		return findIndex(number, nums, start, end)
	} else {
		start = middle + 1
		if start > end {
			return failed
		}
		return findIndex(number, nums, start, end)
	}
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	start, end := 0, length-1
	for ; nums[start]+nums[end] > target; end-- {
	}
	for ; nums[start]+nums[end] < target; start++ {
	}

	for ; ; start++ {
		rest := target - nums[start]
		idx := findIndex(rest, nums, start+1, end)
		if idx != failed {
			return []int{start + 1, idx + 1}
		}
	}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
