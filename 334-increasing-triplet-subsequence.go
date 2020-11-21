package main

func increasingTriplet(nums []int) bool {
	var first, second, single int
	var hasFirst, hasSecond, hasSingle bool
	for i := 0; i < len(nums); i++ {
		if !hasFirst {
			hasFirst = true
			first = nums[i]
			continue
		} else if !hasSecond {
			if nums[i] > first {
				second = nums[i]
				hasSecond = true
			} else {
				first = nums[i]
			}
			continue
		}

		if nums[i] > second {
			return true
		} else if first < nums[i] {
			second = nums[i]
			if hasSingle && second > single && first >= single {
				first = single
				hasSingle = false
			}
		} else if !hasSingle {
			hasSingle = true
			single = nums[i]
		} else if single > nums[i] {
			single = nums[i]
		} else if single < nums[i] && nums[i] <= second {
			first = single
			second = nums[i]
			hasSingle = false
		}
	}
	return false
}
