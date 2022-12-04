package main

func dividePlayers(skill []int) int64 {
	var total int
	players := make(map[int]int)
	for _, val := range skill {
		total += val
		players[val]++
	}
	if total%(len(skill)>>1) != 0 {
		return -1
	}
	team := total / (len(skill) >> 1)
	divide := func() int64 {
		var a int
		for v, _ := range players {
			a = v
			break
		}
		players[a]--
		if players[a] == 0 {
			delete(players, a)
		}
		b := team - a
		if _, ok := players[b]; !ok {
			return -1
		}
		players[b]--
		if players[b] == 0 {
			delete(players, b)
		}
		return int64(a * b)
	}
	var result int64
	for len(players) > 0 {
		chemistry := divide()
		if chemistry == -1 {
			return -1
		}
		result += chemistry
	}
	return result
}
