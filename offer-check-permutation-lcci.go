package main

import "sort"

func CheckPermutation(s1 string, s2 string) bool {
	rs1, rs2 := []rune(s1), []rune(s2)
	sort.Slice(rs1, func(i, j int) bool {
		return rs1[i] < rs1[j]
	})
	sort.Slice(rs2, func(i, j int) bool {
		return rs2[i] < rs2[j]
	})
	return string(rs1) == string(rs2)
}
