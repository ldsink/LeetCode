package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getMax(a, b int, nums []int, max, min map[[2]int]float64) float64 {
	if v, ok := max[[2]int{a, b}]; ok {
		return v
	}
	if a == b {
		max[[2]int{a, b}] = float64(nums[a])
	} else {
		var value float64
		for i := a; i < b; i++ {
			if v := getMax(a, i, nums, max, min) / getMin(i+1, b, nums, max, min); value < v {
				value = v
			}
		}
		max[[2]int{a, b}] = value
	}
	return max[[2]int{a, b}]
}

func getMin(a, b int, nums []int, max, min map[[2]int]float64) float64 {
	if v, ok := min[[2]int{a, b}]; ok {
		return v
	}
	if a == b {
		min[[2]int{a, b}] = float64(nums[a])
	} else {
		value := math.MaxFloat64
		for i := a; i < b; i++ {
			if v := getMin(a, i, nums, max, min) / getMax(i+1, b, nums, max, min); value > v {
				value = v
			}
		}
		min[[2]int{a, b}] = value
	}
	return min[[2]int{a, b}]
}

func getMaxString(a, b int, nums []int, max, min map[[2]int]float64) string {
	if a == b {
		return strconv.Itoa(nums[a])
	} else if a+1 == b {
		return fmt.Sprintf("%d/%d", nums[a], nums[b])
	}
	for i := a; i < b; i++ {
		if v := max[[2]int{a, i}] / min[[2]int{i + 1, b}]; v == max[[2]int{a, b}] {
			numerator := getMaxString(a, i, nums, max, min)
			denominator := getMinString(i+1, b, nums, max, min)
			// 如果分母包含除法，需要加括号
			if strings.Contains(denominator, "/") {
				denominator = fmt.Sprintf("(%s)", denominator)
			}
			return fmt.Sprintf("%s/%s", numerator, denominator)
		}
	}
	return ""
}

func getMinString(a, b int, nums []int, max, min map[[2]int]float64) string {
	if a == b {
		return strconv.Itoa(nums[a])
	} else if a+1 == b {
		return fmt.Sprintf("%d/%d", nums[a], nums[b])
	}
	for i := a; i < b; i++ {
		if v := min[[2]int{a, i}] / max[[2]int{i + 1, b}]; v == min[[2]int{a, b}] {
			numerator := getMinString(a, i, nums, max, min)
			denominator := getMaxString(i+1, b, nums, max, min)
			// 如果分母包含除法，需要加括号
			if strings.Contains(denominator, "/") {
				denominator = fmt.Sprintf("(%s)", denominator)
			}
			return fmt.Sprintf("%s/%s", numerator, denominator)
		}
	}
	return ""
}

func optimalDivision(nums []int) string {
	max := make(map[[2]int]float64)
	min := make(map[[2]int]float64)
	getMax(0, len(nums)-1, nums, max, min)
	return getMaxString(0, len(nums)-1, nums, max, min)
}
