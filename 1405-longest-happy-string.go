package main

func longestDiverseString(a int, b int, c int) string {
	var result []rune
	count := [4]int{0, a, b, c}        // count[0] 用来存上一次的字符
	chars := [4]rune{0, 'a', 'b', 'c'} // chars[0] 用来存上一次字符出现的次数
	for count[1]+count[2]+count[3] > 0 {
		var firstRune, firstNum, secondRune, secondNum int // 找到剩余最多的字符和第二多的字符
		for i := 1; i < 4; i++ {
			if count[i] <= 0 {
				continue
			}
			if firstNum < count[i] {
				secondRune, secondNum = firstRune, firstNum
				firstRune, firstNum = i, count[i]
			} else if secondNum < a {
				secondRune, secondNum = i, count[i]
			}
		}
		if firstNum == 0 { // 没有可用的字符了
			break
		} else if firstRune != count[0] { // 当前剩余最多的字符和之前的不相同，直接添加
			count[firstRune] -= 1
			count[0], chars[0] = firstRune, 1
			result = append(result, chars[firstRune])
		} else if chars[0] == 1 { // 字符只使用了一次
			count[firstRune] -= 1
			chars[0]++
			result = append(result, chars[firstRune])
		} else if secondNum == 0 { // 剩余最多的字符已被连续使用两次，没有其他字符
			break
		} else {
			count[secondRune] -= 1
			count[0], chars[0] = secondRune, 1
			result = append(result, chars[secondRune])
		}
	}
	return string(result)
}
