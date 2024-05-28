package main

func findPeaks(mountain []int) []int {
	var result []int
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i-1] < mountain[i] && mountain[i] > mountain[i+1] {
			result = append(result, i)
			i++
		}
	}
	return result
}
