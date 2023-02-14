package main

func longestWPI(hours []int) int {
	prefixSum := []int{0}
	sumIndex := make(map[int]int)
	for idx, hour := range hours {
		if hour > 8 {
			prefixSum = append(prefixSum, prefixSum[idx]+1)
		} else {
			prefixSum = append(prefixSum, prefixSum[idx]-1)
		}
		if _, ok := sumIndex[prefixSum[idx+1]]; !ok {
			sumIndex[prefixSum[idx+1]] = idx + 1
		}
	}
	var result int
	for i := len(prefixSum) - 1; i > result; i-- {
		if prefixSum[i] > 0 {
			if result < i {
				result = i
			}
		} else if idx, ok := sumIndex[prefixSum[i]-1]; ok && idx < i {
			if result < i-idx {
				result = i - idx
			}
		}
	}
	return result
}
