package main

func getWinner(arr []int, k int) int {
	maxNum := arr[0]
	for _, num := range arr {
		if maxNum < num {
			maxNum = num
		}
	}
	for i, j, win := 0, 1, 0; i < len(arr); {
		if arr[i] == maxNum {
			return maxNum
		}
		if arr[i] >= arr[j] {
			arr = append(arr, arr[j])
			j++
			win++
		} else {
			arr = append(arr, arr[i])
			arr = arr[j:]
			i, j, win = 0, 1, 1
		}
		if win == k {
			return arr[i]
		}
	}
	return -1
}
