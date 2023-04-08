package main

import "math/bits"

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	skillNum := make(map[string]int)
	for idx, skill := range req_skills {
		skillNum[skill] = idx
	}
	var peoples []uint
	for _, skills := range people {
		var ability uint
		for _, skill := range skills {
			ability ^= 1 << skillNum[skill]
		}
		peoples = append(peoples, ability)
	}
	type Team struct {
		people, skill uint
	}
	queue := []Team{{people: 0, skill: 0}}
	smallest := make(map[uint]uint)
	smallest[0] = 0
	for i, s := range peoples {
		var current []Team
		for _, t := range queue {
			if smallest[t.skill] != t.people { // 非最优解，肯定不用扩展继续加入队列
				continue
			}
			current = append(current, t) // 继续加入队列，为下一次做准备
			skill := t.skill | s
			if prevTeam, ok := smallest[skill]; !ok || bits.OnesCount(prevTeam) > bits.OnesCount(t.people)+1 {
				people := t.people ^ (1 << i)
				smallest[skill] = people
				current = append(current, Team{people: people, skill: skill})
			}
		}
		queue = current
	}
	var result []int
	target := uint(1<<len(req_skills) - 1)
	for i := 0; i < len(people); i++ {
		if (1<<i)&smallest[target] != 0 {
			result = append(result, i)
		}
	}
	return result
}
