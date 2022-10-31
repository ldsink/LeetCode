package main

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}
	arr := []bool{true, false, false} // true for 1, zero for 2
	for i, j := 2, true; len(arr) <= n; i, j = i+1, !j {
		arr = append(arr, j)
		if !arr[i] { // 当前元素出现两次
			arr = append(arr, j)
		}
	}
	var count int
	for i := 0; i < n; i++ {
		if arr[i] {
			count++
		}
	}
	return count
}
