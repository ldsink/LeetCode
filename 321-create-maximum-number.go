package main

var array1, array2 []int
var cache map[[3]int][]int

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

func search(index1, index2, k int) {
	key := [3]int{index1, index2, k}
	if _, ok := cache[key]; ok {
		return
	}
	if k == 0 {
		cache[key] = []int{}
		return
	}

	var limit, idx1, idx2 int
	num1, num2 := -1, -1
	// get array1 max number
	limit = len(array1)
	if k-(len(array2)-index2) > 0 {
		limit -= k - (len(array2) - index2) - 1
	}
	for i := index1; i < limit; i++ {
		if num1 < array1[i] {
			num1 = array1[i]
			idx1 = i
		}
		if num1 == 9 {
			break
		}
	}
	// get array2 max number
	limit = len(array2)
	if k-(len(array1)-index1) > 0 {
		limit -= k - (len(array1) - index1) - 1
	}
	for i := index2; i < limit; i++ {
		if num2 < array2[i] {
			num2 = array2[i]
			idx2 = i
		}
		if num2 == 9 {
			break
		}
	}

	nextA := [3]int{idx1 + 1, index2, k - 1}
	nextB := [3]int{index1, idx2 + 1, k - 1}
	if num1 > num2 || index2 == len(array2) {
		search(idx1+1, index2, k-1)
		cache[key] = append([]int{num1}, cache[nextA]...)
		return
	} else if num1 < num2 || index1 == len(array1) {
		search(index1, idx2+1, k-1)
		cache[key] = append([]int{num2}, cache[nextB]...)
		return
	}
	search(idx1+1, index2, k-1)
	search(index1, idx2+1, k-1)
	if isMax(cache[nextA], cache[nextB]) {
		cache[key] = append([]int{num1}, cache[nextA]...)
	} else {
		cache[key] = append([]int{num2}, cache[nextB]...)
	}
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	array1 = nums1
	array2 = nums2
	cache = make(map[[3]int][]int)
	search(0, 0, k)
	return cache[[3]int{0, 0, k}]
}
