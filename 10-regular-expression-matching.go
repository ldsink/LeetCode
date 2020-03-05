package main

const DOT rune = '.'
const WILDCARD rune = '*'

var sr, pr []rune

type regexBlock struct {
	Type   int // 1 a; 2 a*; 3 .; 4 .*
	Letter rune
}

type Mark map[int]map[int]int

func _getRegexBlock(index int) regexBlock {
	r := pr[index]
	if index+1 < len(pr) && pr[index+1] == WILDCARD {
		if r == DOT {
			return regexBlock{Type: 4, Letter: DOT}
		} else {
			return regexBlock{Type: 2, Letter: r}
		}
	}
	if r == DOT {
		return regexBlock{Type: 3, Letter: DOT}
	}
	return regexBlock{Type: 1, Letter: r}
}

func _isEmptyRegex(index int) bool {
	for ; index < len(pr); {
		pb := _getRegexBlock(index)
		if pb.Type == 2 || pb.Type == 4 {
			index += 2
		} else {
			return false
		}
	}
	return true
}

func _setMarkAndReturn(sIndex, pIndex int, markPtr *Mark, success bool) bool {
	status := -1
	if success {
		status = 1
	}
	(*markPtr)[sIndex][pIndex] = status
	return success
}

func _isMatch(sIndex, pIndex int, markPtr *Mark) bool {
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

	success := false
	sEndFlag := len(sr) == sIndex
	pEndFlag := len(pr) == pIndex
	if sEndFlag || pEndFlag {
		if sEndFlag && pEndFlag {
			success = true
		} else if pEndFlag {
			success = false
		} else {
			success = _isEmptyRegex(pIndex)
		}
		return _setMarkAndReturn(sIndex, pIndex, markPtr, success)
	}

	pb := _getRegexBlock(pIndex)
	if pb.Type == 1 {
		if pb.Letter == sr[sIndex] {
			success = _isMatch(sIndex+1, pIndex+1, markPtr)
		} else {
			success = false
		}
	} else if pb.Type == 2 {
		for i := sIndex; i < len(sr) && sr[i] == pb.Letter; i++ {
			if _isMatch(i+1, pIndex+2, markPtr) {
				success = true
				break
			}
		}
		if !success {
			success = _isMatch(sIndex, pIndex+2, markPtr)
		}
	} else if pb.Type == 3 {
		success = _isMatch(sIndex+1, pIndex+1, markPtr)
	} else {
		for i := sIndex; i < len(sr); i++ {
			if _isMatch(i+1, pIndex+2, markPtr) {
				success = true
				break
			}
		}
		if !success {
			success = _isMatch(sIndex, pIndex+2, markPtr)
		}
	}
	return _setMarkAndReturn(sIndex, pIndex, markPtr, success)
}

func isMatch(s string, p string) bool {
	sr = []rune(s)
	pr = []rune(p)
	mark := make(Mark)
	return _isMatch(0, 0, &mark)
}
