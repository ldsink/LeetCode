package main

import "fmt"

func findMiddle(nums []int) float64 {
	length := len(nums)
	if length%2 == 0 {
		median1 := nums[length>>1]
		median2 := nums[length>>1-1]
		return float64(median1+median2) / 2
	}
	return float64(nums[length>>1])
}

func getNthValue(n, l1, r1, l2, r2 int, nums1, nums2 []int) int {
	// TODO
	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMiddle(nums2)
	} else if len(nums2) == 0 {
		return findMiddle(nums1)
	}

	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		n1 := length >> 1
		n2 := n1 - 1
		median1 := getNthValue(n1, 0, len(nums1)-1, 0, len(nums2)-1, nums1, nums2)
		median2 := getNthValue(n2, 0, len(nums1)-1, 0, len(nums2)-1, nums1, nums2)
		return float64(median1+median2) / 2
	}
	median := getNthValue(length>>1, 0, len(nums1)-1, 0, len(nums2)-1, nums1, nums2)
	return float64(median)
}

func main() {
	var nums1, nums2 []int
	var result float64

	nums1 = []int{1, 2, 3}
	nums2 = []int{}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)

	nums1 = []int{}
	nums2 = []int{1, 2, 3, 4}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
