package main

func maxHappyGroups(batchSize int, groups []int) int {
	const bitLen = 5
	const bitMask = 0x1F
	var result, state int
	for _, group := range groups {
		index := group % batchSize
		if index == 0 {
			result++
			continue
		}
		state += 1 << (index * bitLen)
	}
	cached := make(map[int]int)
	var dfs func(int, int) int
	dfs = func(state, donuts int) int {
		if v, ok := cached[state]; ok {
			return v
		}
		var maxGroup, happy int
		if donuts == 0 {
			happy = 1
		}
		for i := 1; i < batchSize; i++ {
			if state>>(i*bitLen)&bitMask != 0 {
				tmp := dfs(state-1<<(i*bitLen), (donuts+i)%batchSize)
				if t := tmp + happy; maxGroup < t {
					maxGroup = t
				}
			}
		}
		cached[state] = maxGroup
		return maxGroup
	}
	return result + dfs(state, 0)
}
