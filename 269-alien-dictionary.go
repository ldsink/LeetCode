package main

import (
	"strings"
)

// https://leetcode.cn/problems/Jf1JuT/
func alienOrder(words []string) string {
	const max = 26
	var parents, children [max]map[int]bool
	for i := 0; i < max; i++ {
		parents[i] = make(map[int]bool)
		children[i] = make(map[int]bool)
	}
	getIdx := func(b rune) int { return int(b - 'a') }

	// 验证部分的正确性，返回去掉公共前缀的结果
	validatePart := func(ss []string) ([]string, bool) {
		var result []string
		for _, s := range ss {
			result = append(result, s[1:])
		}
		// 空字符串必须在顶部，否则不正确
		for i := 0; i < len(result); i++ {
			if result[i] != "" {
				continue
			}
			if i == 0 {
				result = result[1:]
				i--
				continue
			}
			return []string{}, false
		}
		return result, true
	}

	var validate func([]string) bool
	validate = func(ss []string) bool {
		var curr []string
		var success bool
		var order []rune // 本次 ss 第一个字母的字典序
		for _, s := range ss {
			// 当前为空或者是同一组，需要添加
			if len(curr) == 0 || s[0] == curr[0][0] {
				curr = append(curr, s)
				continue
			}
			// 碰到一组新的情况
			order = append(order, rune(curr[0][0]))
			curr, success = validatePart(curr)
			if !success || !validate(curr) {
				return false
			}
			curr = []string{s}
		}
		if len(curr) > 0 {
			order = append(order, rune(curr[0][0]))
			curr, success = validatePart(curr)
			if !success || !validate(curr) {
				return false
			}
		}
		if len(order) > 0 {
			for i := 1; i < len(order); i++ {
				prev, curr := getIdx(order[i-1]), getIdx(order[i])
				if !parents[curr][prev] {
					parents[curr][prev] = true
				}
				if !children[prev][curr] {
					children[prev][curr] = true
				}
				if children[curr][prev] || parents[prev][curr] { // 行成回环，返回 false
					return false
				}
			}
		}
		return true
	}

	if !validate(words) {
		return ""
	}

	// 生成一组合法的顺序
	chars := make(map[rune]bool)
	for _, r := range []rune(strings.Join(words, "")) {
		chars[r] = true
	}
	length := len(chars)
	var result []rune
	for success := true; success; {
		success = false
		for r := range chars {
			idx := getIdx(r)
			if len(parents[idx]) > 0 {
				continue
			}
			result = append(result, r)
			for j, _ := range children[idx] { // 删除子节点的父类
				delete(parents[j], idx)
			}
			delete(chars, r)
			success = true
			break
		}
	}
	if len(result) == length {
		return string(result)
	}
	return ""
}
