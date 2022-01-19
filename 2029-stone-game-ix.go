package main

func aliceWin(firstTake int, count [3]int) bool {
	var value int
	bout := true
	for true {
		if count[0]+count[1]+count[2] == 0 {
			return false // 移除后没有任何剩余的石子，那么 Bob 将会直接获胜（即便是在 Alice 的回合）。
		}
		if value == 0 {
			// 按照 firstTake 取数，取不到都算 alice 都输
			if count[firstTake] == 0 {
				return false
			}
			count[firstTake]--
			value = firstTake
			bout = !bout
			continue
		}
		// 此时 value 一定大于 0，取 3 的整数倍是安全的。如果是偶数次，先都取掉。剩下 0 或者 1 次。
		if count[0] > 1 {
			count[0] %= 2
			continue
		}
		// value 只可能是 1 或者 2，只能取自身或者 3，否则当前回合取石头的玩家输
		if count[value] == 0 && count[0] == 0 {
			return !bout
		}
		if count[value] > 0 {
			count[value]--
			value = (value * 2) % 3
		} else {
			count[0]--
		}
		bout = !bout
	}
	return false
}

func stoneGameIX(stones []int) bool {
	var count [3]int
	for _, stone := range stones {
		count[stone%3]++
	}
	// 如果 Alice 获胜，返回 true ；如果 Bob 获胜，返回 false 。
	if aliceWin(1, [3]int{count[0], count[1], count[2]}) || aliceWin(2, [3]int{count[0], count[1], count[2]}) {
		return true
	}
	return false
}
