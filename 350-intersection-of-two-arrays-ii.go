package main

func intersect(nums1 []int, nums2 []int) []int {
	n1 := make(map[int]int)
	n2 := make(map[int]int)
	for _, n := range nums1 {
		n1[n] += 1
	}
	for _, n := range nums2 {
		n2[n] += 1
	}
	var result []int
	for n, c := range n1 {
		if n2[n] == 0 {
			continue
		}
		if c > n2[n] {
			c = n2[n]
		}
		for i := 0; i < c; i++ {
			result = append(result, n)
		}
	}
	return result
}
