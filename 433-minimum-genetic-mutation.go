package main

const choices = "ACGT"

func search(start, end string, searched map[string]int, gens map[string]bool) int {
	if start == end {
		return 0
	} else if times, ok := searched[start]; ok {
		return times
	}
	searched[start] = -1 // 先初始化成 -1 避免循环搜索
	rs := []rune(start)
	for i := 0; i < 8; i++ {
		origin := rs[i]
		for _, r := range choices {
			if r == origin {
				continue
			}
			rs[i] = r
			if gens[string(rs)] { // 是一个合法的基因组
				if r := search(string(rs), end, searched, gens); r >= 0 { // 可以转换成 end
					if searched[start] == -1 || searched[start] > r+1 { // 可以更新当前结果
						searched[start] = r + 1
					}
				}
			}
		}
		rs[i] = origin
	}
	return searched[start]
}

func minMutation(start string, end string, bank []string) int {
	gens := make(map[string]bool)
	for _, gen := range bank {
		gens[gen] = true
	}
	if !gens[end] {
		return -1
	}
	searched := make(map[string]int)
	return search(start, end, searched, gens)
}
