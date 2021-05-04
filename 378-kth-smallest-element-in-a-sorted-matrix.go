package main

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	idx := make([]int, n)
	var result int
	for i := 0; i < k; i++ {
		pick := -1
		for j := 0; j < n; j++ {
			if idx[j] == n {
				continue
			}
			if pick == -1 || matrix[j][idx[j]] < matrix[pick][idx[pick]] {
				pick = j
				continue
			}
			if idx[j] == 0 {
				break
			}
		}
		result = matrix[pick][idx[pick]]
		idx[pick] ++
	}
	return result
}
