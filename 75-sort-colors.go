package main

import "fmt"

func sortColors(nums []int) {
	var red, white, i int
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			red++
		} else if nums[i] == 1 {
			white++
		}
	}
	fmt.Println(red, white)
	for i = 0; i < red; i++ {
		nums[i] = 0
	}
	for ; i < red+white; i++ {
		nums[i] = 1
	}
	for ; i < len(nums); i++ {
		nums[i] = 2
	}
}
