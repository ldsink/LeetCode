package main

type letters [2]uint // 使用 128 位二进制位来存储情况

func getIdxAndOffset(i int) (int, int) {
	return i / 16, (i % 16) * 4
}

func stickerToLetters(sticker string) letters {
	var count [26]int
	for _, r := range sticker {
		count[r-'a']++
	}
	var result letters
	for i := 0; i < 26; i++ {
		idx, offset := getIdxAndOffset(i)
		result[idx] |= uint(count[i]) << offset
	}
	return result
}

// 从 target 里面减去 sticker 的字符数量
func minusLetters(sticker, target letters) letters {
	var result letters
	for i := 0; i < 26; i++ { // 遍历 26 个字母处理情况
		idx, offset := getIdxAndOffset(i)
		a := (target[idx] >> offset) & 15
		if a > 0 {
			if b := (sticker[idx] >> offset) & 15; b > 0 {
				if a >= b {
					a -= b
				} else {
					a = 0
				}
			}
		}
		result[idx] |= (a & 15) << offset
	}
	return result
}

func dp(stickers []letters, stickerIdx int, target letters, cached map[[3]uint]int) int {
	key := [3]uint{target[0], target[1], uint(stickerIdx)}
	// 已经搜索过，不重复搜索
	if r, ok := cached[key]; ok {
		return r
	}
	// 到贴纸尾部了，还没处理完成
	if stickerIdx == len(stickers) {
		cached[key] = -1
		return -1
	}
	result := dp(stickers, stickerIdx+1, target, cached)
	for i, prev := 1, target; i < 16; i++ { // 每个贴纸最多使用 16 次。
		next := minusLetters(stickers[stickerIdx], prev)
		if next == prev { // 不能带来字符变动，停止计算
			break
		}
		if next[0] == 0 && next[1] == 0 { // 用 i 次贴纸可以完成
			if result == -1 || result > i {
				result = i
			}
			break
		}
		if count := dp(stickers, stickerIdx+1, next, cached); count != -1 {
			if result == -1 || result > count+i {
				result = count + i
			}
		}
		prev = next
	}
	cached[key] = result
	return result
}

func minStickers(stickers []string, target string) int {
	cached := make(map[[3]uint]int)
	var newLetters []letters
	for _, sticker := range stickers {
		newLetters = append(newLetters, stickerToLetters(sticker))
	}
	return dp(newLetters, 0, stickerToLetters(target), cached)
}
