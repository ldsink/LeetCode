package main

import (
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	var result string
	if numerator*denominator < 0 {
		result += "-"
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}

	result += fmt.Sprintf("%d", numerator/denominator)
	numerator = numerator % denominator
	if numerator == 0 {
		return result
	} else {
		result += "."
	}

	var numerators, nums []int
	indexes := make(map[int]int)
	numerators = append(numerators, numerator)
	indexes[numerator] = len(numerators) - 1
	for i := 0; i < len(numerators); i++ {
		numerator = numerators[i] * 10
		nums = append(nums, numerator/denominator)
		numerator = numerator % denominator
		if numerator == 0 {
			for j := 0; j < len(nums); j++ {
				result += fmt.Sprintf("%d", nums[j])
			}
			return result
		}
		if index, ok := indexes[numerator]; ok {
			for j := 0; j < index; j++ {
				result += fmt.Sprintf("%d", nums[j])
			}
			result += "("
			for j := index; j < len(nums); j++ {
				result += fmt.Sprintf("%d", nums[j])
			}
			result += ")"
			return result
		}
		numerators = append(numerators, numerator)
		indexes[numerator] = len(numerators) - 1
	}
	return result
}
