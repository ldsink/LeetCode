package main

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	type Player struct {
		score, age int
	}
	var players []Player
	for i, score := range scores {
		players = append(players, Player{score: score, age: ages[i]})
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].age < players[j].age || (players[i].age == players[j].age && players[i].score < players[j].score)
	})
	var sumScore []int
	var result int
	for i, player := range players {
		var s, c int
		for j := i - 1; j >= c; j-- {
			if player.age > players[j].age && player.score < players[j].score {
				continue
			} else if s < sumScore[j] {
				s = sumScore[j]
			}
		}
		s += player.score
		sumScore = append(sumScore, s)
		if result < s {
			result = s
		}
	}
	return result
}
