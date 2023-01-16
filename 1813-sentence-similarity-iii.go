package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	part1 := strings.Split(sentence1, " ")
	part2 := strings.Split(sentence2, " ")
	if len(part1) < len(part2) {
		part1, part2 = part2, part1
	}
	for len(part2) > 0 {
		if part1[0] == part2[0] {
			part1 = part1[1:]
			part2 = part2[1:]
			continue
		} else if part1[len(part1)-1] == part2[len(part2)-1] {
			part1 = part1[:len(part1)-1]
			part2 = part2[:len(part2)-1]
			continue
		}
		break
	}
	return len(part2) == 0
}
