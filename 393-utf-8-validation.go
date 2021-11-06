package main

import (
	"strconv"
)

func toBinaryString(num int) string {
	result := strconv.FormatInt(int64(num), 2)
	for len(result) < 8 {
		result = "0" + result
	}
	return result[:8]
}

func validUtf8(data []int) bool {
	if len(data) == 0 {
		return true
	}
	s := toBinaryString(data[0])
	if s[:1] == "0" {
		return validUtf8(data[1:])
	}
	prefixes := []string{"110", "1110", "11110"}
	for index, prefix := range prefixes {
		if s[:len(prefix)] != prefix {
			continue
		}
		if len(data) < index+2 {
			return false
		}
		for i := 0; i <= index; i++ {
			c := toBinaryString(data[1+i])
			if c[:2] != "10" {
				return false
			}
		}
		return validUtf8(data[index+2:])
	}
	return false
}
