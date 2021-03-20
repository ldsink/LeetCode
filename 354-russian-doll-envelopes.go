package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] == envelopes[j][1])
	})
	depth := 1
	dolls := make(map[int][][]int)
	for _, envelope := range envelopes {
		added := false
		for d := depth; d > 0 && !added; d-- {
			for _, child := range dolls[d] {
				if child[0] < envelope[0] && child[1] < envelope[1] {
					added = true
					dolls[d+1] = append(dolls[d+1], envelope)
					if depth < d+1 {
						depth = d + 1
					}
					break
				}
			}
		}
		dolls[1] = append(dolls[1], envelope)
	}
	return depth
}
