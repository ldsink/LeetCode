package main

import (
	"fmt"
	"math"
)

var length, target int
var median1, median2 float64

func binarySearch(primary []int, secondary []int) bool {
	for i, j := 0, len(primary)-1; i <= j; {
		middle := (i + j) >> 1
		rest := target - (middle + 1) - 1
		if rest+1 < 0 || (rest+1 < len(secondary) && primary[middle] > secondary[rest+1]) {
			j = middle - 1
		} else if rest >= len(secondary) || (rest >= 0 && primary[middle] < secondary[rest]) {
			i = middle + 1
		} else {
			median2 = float64(primary[middle])
			if length%2 != 0 {
				median1 = median2
			} else {
				flag := false
				if rest >= 0 {
					median1 = float64(secondary[rest])
					flag = true
				}
				if middle >= 1 {
					if flag {
						median1 = math.Max(median1, float64(primary[middle-1]))
					} else {
						median1 = float64(primary[middle-1])
					}
				}
			}
			return true
		}
	}
	return false
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length = len(nums1) + len(nums2)
	target = (length)>>1 + 1

	if !binarySearch(nums1, nums2) {
		binarySearch(nums2, nums1)
	}

	return (median1 + median2) / 2
}

func main() {
	var nums1, nums2 []int
	var result float64

	nums1 = []int{1}
	nums2 = []int{}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)

	nums1 = []int{1, 2, 3}
	nums2 = []int{}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)

	nums1 = []int{2}
	nums2 = []int{1, 3, 4}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)

	nums1 = []int{1, 3}
	nums2 = []int{2}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)

	nums1 = []int{1, 7, 8}
	nums2 = []int{2, 3, 4, 5, 10, 30}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println("result: ", result)
}
