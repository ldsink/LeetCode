package main

import "fmt"

func maxProduct(nums []int) int {
	result := nums[0]
	var minusCount int
	start, firstMinusIndex, lastMinusIndex := -1, -1, -1
	for index, num := range nums {
		if start == -1 {
			start = index
		}
		if num < 0 {
			minusCount++
			lastMinusIndex = index
			if firstMinusIndex == -1 {
				firstMinusIndex = index
			}
		} else if num == 0 {
			if result < 0 {
				result = 0
			}
			if start == index {
				start = -1
				continue
			} else if minusCount%2 == 0 {
				product := 1
				for i := start; i < index; i++ {
					product *= nums[i]
				}
				if result < product {
					result = product
				}
			} else if firstMinusIndex == lastMinusIndex {
				if start < firstMinusIndex {
					product := 1
					for i := start; i < firstMinusIndex; i++ {
						product *= nums[i]
					}
					if result < product {
						result = product
					}
				}
				if firstMinusIndex+1 < index {
					product := 1
					for i := firstMinusIndex + 1; i < index; i++ {
						product *= nums[i]
					}
					if result < product {
						result = product
					}
				}
				if result < nums[firstMinusIndex] {
					result = nums[firstMinusIndex]
				}
			} else {
				middle := 1
				for i := firstMinusIndex + 1; i < lastMinusIndex; i++ {
					middle *= nums[i]
				}
				product := nums[firstMinusIndex] * middle
				for i := start; i < firstMinusIndex; i++ {
					product *= nums[i]
				}
				if result < product {
					result = product
				}
				product = middle * nums[lastMinusIndex]
				for i := lastMinusIndex + 1; i < index; i++ {
					product *= nums[i]
				}
				if result < product {
					result = product
				}
			}
			start, firstMinusIndex, lastMinusIndex = -1, -1, -1
			minusCount = 0
		}
	}

	// 处理最后的结尾
	if start != -1 {
		if minusCount%2 == 0 {
			product := 1
			for i := start; i < len(nums); i++ {
				product *= nums[i]
			}
			if result < product {
				result = product
			}
		} else if firstMinusIndex == lastMinusIndex {
			if start < firstMinusIndex {
				product := 1
				for i := start; i < firstMinusIndex; i++ {
					product *= nums[i]
				}
				if result < product {
					result = product
				}
			}
			if firstMinusIndex+1 < len(nums) {
				product := 1
				for i := firstMinusIndex + 1; i < len(nums); i++ {
					product *= nums[i]
				}
				if result < product {
					result = product
				}
			}
			if result < nums[firstMinusIndex] {
				result = nums[firstMinusIndex]
			}
		} else {
			middle := 1
			for i := firstMinusIndex + 1; i < lastMinusIndex; i++ {
				middle *= nums[i]
			}
			product := nums[firstMinusIndex] * middle
			for i := start; i < firstMinusIndex; i++ {
				product *= nums[i]
			}
			if result < product {
				result = product
			}
			product = middle * nums[lastMinusIndex]
			for i := lastMinusIndex + 1; i < len(nums); i++ {
				product *= nums[i]
			}
			if result < product {
				result = product
			}
		}
	}
	return result
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, 2, -2, -3, -5}))
}
