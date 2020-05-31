package main

import (
	"strconv"
)

func dfs(target int, result *[]string, num, prefix string, value, delta int) {
	if len(num) == 0 {
		if target == value {
			*result = append(*result, prefix)
		}
		return
	}
	for i := 1; i <= len(num); i++ {
		leftNum := num[:i]
		if len(leftNum) > 1 && leftNum[0] == '0' {
			break
		}
		n, _ := strconv.Atoi(leftNum)
		if len(prefix) > 0 {
			dfs(target, result, num[i:], prefix+"+"+leftNum, value+n, n)
			dfs(target, result, num[i:], prefix+"-"+leftNum, value-n, -n)
			dfs(target, result, num[i:], prefix+"*"+leftNum, value-delta+(delta*n), delta*n)
		} else {
			dfs(target, result, num[i:], leftNum, n, n)
		}
	}
}

func addOperators(num string, target int) []string {
	var result []string
	dfs(target, &result, num, "", 0, 0)
	return result
}
