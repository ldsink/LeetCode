package main

func majorityElement(nums []int) []int {
	var num1, num2, count1, count2 int
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		} else if count1 == 0 {
			num1 = num
			count1++
		} else if count2 == 0 {
			num2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		}
	}

	var result []int
	if count1 > len(nums)/3 {
		result = append(result, num1)
	}
	if count2 > len(nums)/3 {
		result = append(result, num2)
	}
	return result
}
