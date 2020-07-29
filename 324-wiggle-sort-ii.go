package main

func isFit(nums []int, idx, num int) bool {
	if idx&1 == 0 {
		return num < nums[idx+1] && (idx == 0 || num < nums[idx-1])
	} else {
		return num > nums[idx+1] && num > nums[idx-1]
	}
}

func findAfterDiffNum(nums []int, idx int) int {
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			return i
		}
	}
	return idx
}

func findBeforeDiffNum(nums []int, idx int) int {
	for j := idx - 1; j >= 0; j-- {
		if nums[j] != nums[idx] && isFit(nums, j, nums[idx]) {
			return j
		}
	}
	return idx
}

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i&1 == 0 {
			// 低点的情况
			if nums[i] < nums[i-1] {
				continue
			} else if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				continue
			}
			// 相等的情况，从前后寻找一个不同的数替换这个数。
			j := findAfterDiffNum(nums, i)
			if j == i {
				j = findBeforeDiffNum(nums, i)
			}
			nums[i], nums[j] = nums[j], nums[i]
			i--
			continue
		} else {
			// 高点的情况
			if nums[i] > nums[i-1] {
				continue
			} else if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				continue
			}
			// 相等的情况，从前后寻找一个不同的数替换这个数。
			j := findAfterDiffNum(nums, i)
			if j == i {
				j = findBeforeDiffNum(nums, i)
			}
			nums[i], nums[j] = nums[j], nums[i]
			i--
			continue
		}
	}
}
