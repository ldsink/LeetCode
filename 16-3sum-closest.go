package main

import (
	"fmt"
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestInRange(left, right, n int, numsPtr *[]int) int {
	if left == right {
		return (*numsPtr)[left]
	}
	if n >= (*numsPtr)[right] {
		return (*numsPtr)[right]
	}
	if n <= (*numsPtr)[left] {
		return (*numsPtr)[left]
	}

	middle := (left + right) >> 1
	if (*numsPtr)[middle] > n {
		return findClosestInRange(left, middle, n, numsPtr)
	} else if (*numsPtr)[middle] < n {
		if Abs((*numsPtr)[middle+1]-n) > Abs((*numsPtr)[middle]-n) {
			return (*numsPtr)[middle]
		}
		return findClosestInRange(middle+1, right, n, numsPtr)
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	offset := Abs(result - target)
	for i := 0; i < len(nums)-2; i++ {
		for j := len(nums) - 1; j > i+1; j-- {
			k := target - nums[i] - nums[j]
			l := findClosestInRange(i+1, j-1, k, &nums)
			r := nums[i] + nums[j] + l
			o := Abs(r - target)
			if offset > o {
				result = r
				offset = o
			}
		}
	}
	return result
}

func main() {
	var nums []int
	var target int
	nums = []int{-1, 2, 1, -4}
	target = 1
	fmt.Println(threeSumClosest(nums, target))
}
