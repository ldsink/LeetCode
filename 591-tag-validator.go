package main

import (
	"regexp"
	"strings"
)

func isValid(code string) bool {
	var tags []string
	tagRe := regexp.MustCompile("^[A-Z]{1,9}$")
	for len(code) > 0 {
		if code[0] != '<' {
			if len(tags) == 0 { // 不在标签内
				return false
			}
			code = code[1:]
			continue
		}
		if len(code) < 2 { // 确保长度大于 1，这样后面的检查不用考虑越界
			return false
		}
		if code[1] == '/' { // 结束标签的情况
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			if len(tags) == 0 || tags[len(tags)-1] != code[2:j] {
				return false
			}
			tags = tags[:len(tags)-1]
			code = code[j+1:]
			if len(tags) == 0 && len(code) > 0 { // 代码必须被合法的闭合标签包围
				return false
			}
		} else if code[1] == '!' { // cdata 的情况
			if len(tags) == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
				return false
			}
			j := strings.Index(code, "]]>")
			if j == -1 {
				return false
			}
			code = code[j+1:]
		} else { // 标签开始的情况
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			tagName := code[1:j]
			// 合法的 TAG_NAME 仅含有大写字母，长度在范围 [1,9] 之间
			if !tagRe.MatchString(tagName) {
				return false
			}
			tags = append(tags, tagName)
			code = code[j+1:]
		}
	}
	return len(tags) == 0
}
