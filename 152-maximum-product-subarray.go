package main

func maxProductInRange(nums []int, start, end, minusCount, first, last int) int {
	result := nums[start]
	if minusCount%2 == 0 {
		product := 1
		for i := start; i < end; i++ {
			product *= nums[i]
		}
		return product
	}
	// single minus
	if first == last {
		if start < first {
			product := 1
			for i := start; i < first; i++ {
				product *= nums[i]
			}
			if result < product {
				result = product
			}
		}
		if first+1 < end {
			product := 1
			for i := first + 1; i < end; i++ {
				product *= nums[i]
			}
			if result < product {
				result = product
			}
		}
		if result < nums[first] {
			result = nums[first]
		}
		return result
	}
	// odd minus
	middle := 1
	for i := first + 1; i < last; i++ {
		middle *= nums[i]
	}
	product := nums[first] * middle
	for i := start; i < first; i++ {
		product *= nums[i]
	}
	if result < product {
		result = product
	}
	product = middle * nums[last]
	for i := last + 1; i < end; i++ {
		product *= nums[i]
	}
	if result < product {
		result = product
	}
	return result
}

func maxProduct(nums []int) int {
	result := nums[0]
	start, minusCount, first, last := -1, 0, -1, -1
	for index, num := range nums {
		if start == -1 {
			start = index
		}
		if num < 0 {
			minusCount++
			last = index
			if first == -1 {
				first = index
			}
		} else if num == 0 {
			if result < 0 {
				result = 0
			}
			if start == index {
				start = -1
				continue
			}
			if product := maxProductInRange(nums, start, index, minusCount, first, last); result < product {
				result = product
			}
			start, minusCount, first, last = -1, 0, -1, -1
		}
	}

	if start != -1 {
		if product := maxProductInRange(nums, start, len(nums), minusCount, first, last); result < product {
			result = product
		}
	}
	return result
}
