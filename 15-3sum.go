package main

import (
	"fmt"
	"sort"
)

func binarySearch(left, right, n int, numsPtr *[]int) bool {
	if left == right {
		return (*numsPtr)[left] == n
	}
	middle := (left + right) >> 1
	if (*numsPtr)[middle] > n {
		return binarySearch(left, middle, n, numsPtr)
	} else if (*numsPtr)[middle] < n {
		return binarySearch(middle+1, right, n, numsPtr)
	}
	return true // (*numsPtr)[middle] == n
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	added := make(map[[2]int]bool)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := len(nums) - 1; j > i+1; j-- {
			k := 0 - nums[i] - nums[j]
			if k < nums[i] {
				continue
			}
			if nums[j] < k {
				break
			}
			if _, ok := added[[2]int{nums[i], k}]; ok {
				continue
			}
			r := binarySearch(i+1, j-1, k, &nums)
			added[[2]int{nums[i], k}] = r
			if r {
				result = append(result, []int{nums[i], k, nums[j]})
			}
		}
	}
	return result
}

func main() {
	var nums []int
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSum(nums))
}
