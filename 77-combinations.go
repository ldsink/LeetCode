package main

func _search(n, k, p int, array *[]bool) [][]int {
	if k > n {
		return [][]int{}
	}
	used := *array
	var combination []int
	if k == 0 {
		for i := 1; i <= n; i++ {
			if used[i] {
				combination = append(combination, i)
			}
		}
		return [][]int{combination}
	}
	if k+p == n {
		var combination []int
		for i := 0; i < p+1; i++ {
			if used[i] {
				combination = append(combination, i)
			}
		}
		for i := p + 1; i <= n; i++ {
			combination = append(combination, i)
		}
		return [][]int{combination}
	}

	var result [][]int
	for i, j := p, n-k; i <= j; i++ {
		used[i+1] = true
		r := _search(n, k-1, i+1, array)
		result = append(result, r...)
		used[i+1] = false
	}
	return result
}

func combine(n int, k int) [][]int {
	used := make([]bool, n+1)
	return _search(n, k, 0, &used)
}
