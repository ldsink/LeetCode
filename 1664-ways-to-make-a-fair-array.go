package main

func waysToMakeFair(nums []int) int {
	var oddSum, evenSum, currOddSum, currEvenSum, result int
	for i, num := range nums {
		if i&1 == 0 {
			evenSum += num
		} else {
			oddSum += num
		}
	}
	for i, num := range nums {
		if i&1 == 0 { // even
			evenSum -= num
			if currEvenSum+oddSum == currOddSum+evenSum {
				result++
			}
			currEvenSum += num
		} else {
			oddSum -= num
			if currEvenSum+oddSum == currOddSum+evenSum {
				result++
			}
			currOddSum += num
		}
	}
	return result
}
