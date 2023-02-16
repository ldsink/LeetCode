package main

import (
	"strconv"
	"strings"
)

func _isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func _isFloat(s string) bool {
	nums := strings.Split(s, ".")
	if len(nums) > 2 {
		return false
	}
	if len(nums) == 2 {
		if len(strings.TrimSpace(nums[0])) != len(nums[0]) {
			return false
		}
		if len(strings.TrimSpace(nums[1])) != len(nums[1]) {
			return false
		}
		if len(nums[1]) > 0 {
			if strings.HasPrefix(nums[1], "+") || strings.HasPrefix(nums[1], "-") {
				return false
			}
		}
		_beforeEmpty := len(nums[0]) == 0 || nums[0] == "+" || nums[0] == "-"
		_before := _isInt(nums[0])
		_after := _isInt(nums[1])
		return (_beforeEmpty && _after) || (_before && _after) || (_before && len(nums[1]) == 0)
	}
	return _isInt(nums[0])
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	nums := strings.Split(s, "e")
	if len(nums) > 2 {
		return false
	} else if len(nums) == 2 {
		return _isFloat(nums[0]) && _isInt(nums[1])
	}
	return _isFloat(nums[0])
}
