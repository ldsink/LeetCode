package main

var sr, pr []rune

func _isMatch(sIndex, pIndex int, markPtr *map[int]map[int]int) bool {
	mark := *markPtr
	if sCache, ok := mark[sIndex]; ok {
		if status, ok := sCache[pIndex]; ok {
			if status == 1 {
				return true
			} else if status == -1 {
				return false
			}
		}
	} else {
		mark[sIndex] = make(map[int]int)
	}

	if sIndex == len(sr) || pIndex == len(pr) {
		if sIndex == len(sr) && pIndex == len(pr) {
			mark[sIndex][pIndex] = 1
			return true
		}
		if pIndex == len(pr) {
			mark[sIndex][pIndex] = -1
			return false
		}
		// sIndex == len(sr)
		// return true when rest of pr is *
		for _, r := range pr[pIndex:] {
			if r != '*' {
				mark[sIndex][pIndex] = -1
				return false
			}
		}
		mark[sIndex][pIndex] = 1
		return true
	}

	result := false
	if pr[pIndex] == '?' || sr[sIndex] == pr[pIndex] {
		result = _isMatch(sIndex+1, pIndex+1, markPtr)
	} else if pr[pIndex] == '*' {
		for i := sIndex; i <= len(sr); i++ {
			r := _isMatch(i, pIndex+1, markPtr)
			if r {
				result = r
				break
			}
		}
	}

	if result {
		mark[sIndex][pIndex] = 1
	} else {
		mark[sIndex][pIndex] = -1
	}
	return result
}

func isMatch(s string, p string) bool {
	sr = []rune(s)
	pr = []rune(p)
	mark := make(map[int]map[int]int)
	return _isMatch(0, 0, &mark)
}
