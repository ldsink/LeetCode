package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	var result string
	a, b := strs[0], strs[1]
	for i := 0; i < Min(len(a), len(b)); i++ {
		if a[i] == b[i] {
			result = a[0 : i+1]
		} else {
			break
		}
	}

	for i := 2; i < len(strs); i++ {
		k := Min(len(strs[i]), len(result))
		if k == 0 {
			return ""
		}
		result = result[0:k]
		for j := 0; j < k; j++ {
			if strs[i][j] != result[j] {
				result = result[0:j]
				break
			}
		}
	}

	return result
}

func main() {
	var strs []string

	strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
