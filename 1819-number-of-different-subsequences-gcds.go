package main

func countDifferentSubsequenceGCDs(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}

	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	maxVal += 1
	var result int
	used := make([]bool, maxVal)
	for _, num := range nums {
		if !used[num] {
			used[num] = true
			result++
		}
	}
	for i := 1; i < maxVal; i++ {
		if used[i] {
			continue
		}
		prevGCD := 0
		for j := i << 1; j < maxVal; j += i {
			if !used[j] {
				continue
			}
			if prevGCD == 0 {
				prevGCD = j
			} else {
				prevGCD = gcd(j, prevGCD)
			}
			if prevGCD == i {
				result++
				break
			}
		}
	}
	return result
}
