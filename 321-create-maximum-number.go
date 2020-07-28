package main

func isMax(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1); i++ {
		if nums1[i] > nums2[i] {
			return true
		} else if nums1[i] < nums2[i] {
			return false
		}
	}
	return false
}

func greedyPrune(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	} else if len(nums) <= k {
		return nums
	}
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := len(result) - 1; j >= 0 && i+(k-j) <= len(nums); j-- {
			if result[j] >= nums[i] {
				break
			}
			result = result[:j]
		}
		if len(result) < k {
			result = append(result, nums[i])
		}
	}
	return result
}

func greedyMerge(nums1, nums2 []int) []int {
	var result []int
	for i, j := 0, 0; i <= len(nums1) && j <= len(nums2); {
		if i == len(nums1) {
			return append(result, nums2[j:]...)
		} else if j == len(nums2) {
			return append(result, nums1[i:]...)
		} else if nums1[i] > nums2[j] {
			result = append(result, nums1[i])
			i++
		} else if nums1[i] < nums2[j] {
			result = append(result, nums2[j])
			j++
		} else {
			l := 1
			for ; i+l < len(nums1) && j+l < len(nums2) && nums1[i+l] == nums2[j+l]; l++ {
			}
			useNum1 := true
			if i+l == len(nums1) {
				useNum1 = false
			} else if j+l == len(nums2) {
				useNum1 = true
			} else if nums1[i+l] < nums2[j+l] {
				useNum1 = false
			}
			if useNum1 {
				result = append(result, nums1[i])
				i++
			} else {
				result = append(result, nums2[j])
				j++
			}
		}
	}
	return result
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	nums1 = greedyPrune(nums1, k)
	nums2 = greedyPrune(nums2, k)
	result := append(nums1, nums2...)[:k]
	var idx1 int
	if k > len(nums2) {
		idx1 = k - len(nums2)
	}
	for i := idx1; i <= k && i <= len(nums1); i++ {
		current := greedyMerge(greedyPrune(nums1, i), greedyPrune(nums2, k-i))
		if isMax(current, result) {
			result = current
		}
	}
	return result
}
