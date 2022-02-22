package main

func pushDominoes(dominoes string) string {
	rs := []rune(dominoes)
	LEFT, DOT, RIGHT := 'L', '.', 'R'
	pushed := make([]bool, len(dominoes)) // 某个位置上的骨牌处理过没
	changed := true
	for changed {
		changed = false
		status := make(map[int][]rune) // 某个位置上的多米诺被推的情况
		for idx, r := range rs {       // 寻找可能的变动点
			if r != DOT && !pushed[idx] {
				pushed[idx] = true
				if r == RIGHT && idx+1 < len(rs) && rs[idx+1] == DOT {
					status[idx+1] = append(status[idx+1], RIGHT)
				} else if r == LEFT && idx-1 >= 0 && rs[idx-1] == DOT {
					status[idx-1] = append(status[idx-1], LEFT)
				}
			}
		}
		for idx, actions := range status {
			if len(actions) == 2 { // 两个方向的力度，不变化
				continue
			}
			rs[idx], changed = actions[0], true // 变化成目标的样子，并标记有变动
		}
	}
	return string(rs)
}
