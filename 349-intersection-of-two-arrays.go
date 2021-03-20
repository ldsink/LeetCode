package main

func intersection(nums1 []int, nums2 []int) []int {
	n1 := make(map[int]bool)
	n2 := make(map[int]bool)
	for _, n := range nums1 {
		n1[n] = true
	}
	for _, n := range nums2 {
		n2[n] = true
	}
	var result []int
	for n, _ := range n1 {
		if n2[n] {
			result = append(result, n)
		}
	}
	return result
}
