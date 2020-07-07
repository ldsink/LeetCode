package main

const leftParentheses rune = '('
const rightParentheses rune = ')'

type condition struct {
	r     rune
	count int
	used  int
	delta int
}

func getCondition(s string) *condition {
	if len(s) == 0 {
		return nil
	}
	cond := condition{r: rune(s[0]), count: 1}
	for i := 1; i < len(s); i++ {
		if rune(s[i]) != cond.r {
			break
		}
		cond.count++
	}
	if cond.r != '(' && cond.r != ')' {
		cond.used = cond.count
	}
	return &cond
}

func search(index, count, left int, data []*condition, result map[string]bool) {
	if index == len(data) {
		if count != 0 || left != 0 {
			return
		}
		// 添加当前结果到 result
		var rs []rune
		for i := 0; i < len(data); i++ {
			if data[i].used > 0 {
				for j := 0; j < data[i].used; j++ {
					rs = append(rs, data[i].r)
				}
			}
		}
		result[string(rs)] = true
		return
	}
	// 其他符号，直接跳过
	if data[index].r != leftParentheses && data[index].r != rightParentheses {
		search(index+1, count, left, data, result)
		return
	}
	// 左括号
	if data[index].r == leftParentheses {
		for i := data[index].count; i >= 0 && count >= data[index].count-i; i-- {
			restCount := count - (data[index].count - i)
			// 剩余的怎么补左括号都多了
			if left+i > data[index].delta+restCount {
				continue
			}
			data[index].used = i
			search(index+1, restCount, left+i, data, result)
		}
		return
	}
	// 右括号
	for i := data[index].count; i >= 0 && count >= data[index].count-i; i-- {
		// 左括号数量少了
		if i > left {
			continue
		}
		restCount := count - (data[index].count - i)
		data[index].used = i
		search(index+1, restCount, left-i, data, result)
	}
}

func removeInvalidParentheses(s string) []string {
	// 计算多余的括号
	var stack []rune
	for _, r := range []rune(s) {
		if r != leftParentheses && r != rightParentheses {
			continue
		}
		if r == rightParentheses && len(stack) > 0 && stack[len(stack)-1] == leftParentheses {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	if len(stack) == 0 {
		return []string{s}
	}

	// 字符串处理成 Condition 结构
	// 首先按照不同的 Rune 聚合
	var zipped []*condition
	for i := 0; i < len(s); {
		c := getCondition(s[i:])
		if c == nil {
			break
		}
		zipped = append(zipped, c)
		i += c.count
	}
	// 更新每个节点上右括号和左括号的差值
	var leftCount, rightCount int
	if zipped[len(zipped)-1].r == leftParentheses {
		leftCount = zipped[len(zipped)-1].count
	} else if zipped[len(zipped)-1].r == rightParentheses {
		rightCount = zipped[len(zipped)-1].count
	}
	for i := len(zipped) - 2; i >= 0; i-- {
		zipped[i].delta = rightCount - leftCount
		if zipped[i].r == leftParentheses {
			leftCount += zipped[i].count
		} else if zipped[i].r == rightParentheses {
			rightCount += zipped[i].count
		}
	}

	// 存储的结果
	result := make(map[string]bool)
	search(0, len(stack), 0, zipped, result)
	var ret []string
	for key, _ := range result {
		ret = append(ret, key)
	}
	return ret
}
