package main

import (
	"fmt"
)

type Number struct {
	value int
	index int
}

func mergeSort(array, sorted []Number, left, right int) {
	if left >= right {
		return
	}
	middle := (left + right) >> 1
	mergeSort(array, sorted, left, middle)
	mergeSort(array, sorted, middle+1, right)

	leftStart, rightStart := left, middle+1
	for position := left; leftStart <= middle || rightStart <= right; position++ {
		if leftStart <= middle && rightStart <= right {
			if array[leftStart].value > array[rightStart].value {
				sorted[position] = array[rightStart]
				rightStart++
			} else {
				sorted[position] = array[leftStart]
				leftStart++
			}
		} else if leftStart <= middle {
			sorted[position] = array[leftStart]
			leftStart++
		} else {
			sorted[position] = array[rightStart]
			rightStart++
		}
	}

	for position := left; position <= right; position++ {
		array[position] = sorted[position]
	}
}

const failed = -1

func findNumber(number int, nums []Number, start, end int) int {
	if start == end {
		if nums[start].value == number {
			return start
		}
		return failed
	}
	middle := (start + end) >> 1
	if nums[middle].value == number {
		return middle
	} else if nums[middle].value > number {
		end = middle - 1
		if end < start {
			return failed
		}
		return findNumber(number, nums, start, end)
	} else {
		start = middle + 1
		if start > end {
			return failed
		}
		return findNumber(number, nums, start, end)
	}
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	array := make([]Number, length)
	sorted := make([]Number, length)
	for i, _ := range array {
		array[i] = Number{index: i, value: nums[i]}
	}

	mergeSort(array, sorted, 0, length-1)

	start, end := 0, length-1
	for ; array[start].value+array[end].value > target; end-- {
	}
	for ; array[start].value+array[end].value < target; start++ {
	}

	for ; ; start++ {
		rest := target - array[start].value
		idx := findNumber(rest, array, start+1, end)
		if idx != failed {
			end = idx
			break
		}
	}
	if array[start].index > array[end].index {
		return []int{array[end].index, array[start].index}
	}
	return []int{array[start].index, array[end].index}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 26
	result := twoSum(nums, target)
	fmt.Println(result)
}
