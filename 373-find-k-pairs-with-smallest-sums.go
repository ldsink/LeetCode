package main

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	startIndex := make([]int, len(nums1))
	maxSum := nums1[len(nums1)-1] + nums2[len(nums2)-1] + 1
	for i := 0; i < k; i++ {
		// init sum
		current := maxSum
		first := 0
		for ; first < len(nums1) && current == maxSum; first++ {
			if startIndex[first] == len(nums2) {
				continue
			}
			if current > nums1[first]+nums2[startIndex[first]] {
				current = nums1[first] + nums2[startIndex[first]]
				break
			}
		}
		if current == maxSum {
			break
		}
		// greedy
		for better := first + 1; better < len(nums1); better++ {
			for j := startIndex[better]; j < startIndex[first]; j++ {
				if current > nums1[better]+nums2[j] {
					current = nums1[better] + nums2[j]
					first = better
					break
				}
			}
		}
		// add pair
		result = append(result, []int{nums1[first], nums2[startIndex[first]]})
		startIndex[first] ++
	}
	return result
}
