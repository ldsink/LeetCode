package main

import (
	"strings"
)

func getFullPath(list []string) []string {
	var result []string
	for i := 0; i < len(list); i++ {
		// 是文件的情况
		if strings.Contains(list[i], ".") {
			result = append(result, list[i])
			continue
		}
		// 是文件夹的情况
		var j int
		var subList []string
		for j = i + 1; j < len(list) && strings.HasPrefix(list[j], "\t"); j++ {
			subList = append(subList, list[j][1:])
		}
		for _, subPath := range getFullPath(subList) {
			result = append(result, list[i]+"/"+subPath)
		}
		i += j - i - 1
	}
	return result
}

func lengthLongestPath(input string) int {
	list := strings.Split(input, "\n")
	list = getFullPath(list)
	var result int
	for _, p := range list {
		if !strings.Contains(p, ".") {
			continue
		}
		if result < len(p) {
			result = len(p)
		}
	}
	return result
}
