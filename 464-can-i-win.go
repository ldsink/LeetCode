package main

func mustWin(choices, max, target int, cache map[int]bool) bool {
	key := (target << 32) | choices
	if r, ok := cache[key]; ok {
		return r
	}
	for i := 1; i <= max; i++ {
		if choices&(1<<i) == 0 {
			continue
		}
		if i >= target || !mustWin(choices^(1<<i), max, target-i, cache) {
			cache[key] = true
			return true
		}
	}
	cache[key] = false
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	choices := 1<<(maxChoosableInteger+1) - 1
	cache := make(map[int]bool)
	return mustWin(choices, maxChoosableInteger, desiredTotal, cache)
}
