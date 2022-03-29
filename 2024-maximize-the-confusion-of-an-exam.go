package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	// 先压缩 answerKey 的数据
	var answers []int
	for head, tail := 0, 0; head < len(answerKey); head = tail {
		for tail = tail + 1; tail < len(answerKey); tail++ {
			if answerKey[tail] != answerKey[head] {
				break
			}
		}
		answers = append(answers, tail-head)
	}
	// 从前往后，处理一遍，获取最长的
	var result, changed, unchanged int
	for head, tail := 0, 0; head < len(answers); head++ {
		current := head & 1
		// 首先更新 tail 的位置，如果超过了，先回退
		if changed > k {
			for tail--; tail > head; tail-- {
				if tail&1 == current {
					unchanged -= answers[tail]
					continue
				}
				changed -= answers[tail]
				if changed <= k {
					break
				}
			}
		}
		// 回退之后，可以继续往后扩展
		for ; tail < len(answers); tail++ {
			if tail&1 == current {
				unchanged += answers[tail]
				continue
			}
			if changed+answers[tail] <= k {
				changed += answers[tail]
			} else {
				break
			}
		}
		// 基于当前的题目，更新最大值
		unused := k - changed
		// 往后扩展，看能不能再多改几个
		if unused > 0 && tail < len(answers) {
			if answers[tail] >= unused {
				unused = 0
			} else {
				unused -= answers[tail]
			}
		}
		// 往前扩展，看能不能再多改几个
		if unused > 0 && head > 0 {
			if answers[head-1] >= unused {
				unused = 0
			} else {
				unused -= answers[head-1]
			}
		}
		if result < unchanged+k-unused {
			result = unchanged + k - unused
		}
		// 更新下一轮的 changed 和 unchanged
		changed, unchanged = unchanged-answers[head], changed
	}
	return result
}
