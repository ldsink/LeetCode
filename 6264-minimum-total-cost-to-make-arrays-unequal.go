package main

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	var result int64
	// 寻找相等的位置中是否有超过一半的众数
	var sameCount, mode, modeCount int
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[i] {
			if modeCount == 0 {
				mode = nums1[i]
			}
			if nums1[i] == mode {
				modeCount++
			} else {
				modeCount--
			}
			sameCount++
			result += int64(i)
		}
	}
	modeCount = 0
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[i] && nums1[i] == mode {
			modeCount++
		}
	}
	// 众数数量不超过需要交换的数量的一半，此时就是结果
	if modeCount <= (sameCount >> 1) {
		return result
	}
	// 否则需要从低到高引入外部的数字参与交换
	count := (modeCount << 1) - sameCount
	for i := 0; i < len(nums1) && count > 0; i++ {
		if nums1[i] != nums2[i] && nums1[i] != mode && nums2[i] != mode {
			result += int64(i)
			count--
		}
	}
	if count != 0 {
		return -1
	}
	return result
}
