package main

func avoidFlood(rains []int) []int {
	full := make(map[int]bool)
	pump := make(map[int]int)
	var result []int
	for i, lake := range rains {
		if lake == 0 {
			result = append(result, -1)
			continue
		}
		if full[lake] {
			for j := pump[lake] + 1; j < i; j++ {
				if result[j] == -1 && rains[j] == 0 {
					result[j] = lake
					delete(full, lake)
					break
				}
			}
		}
		if full[lake] {
			return []int{}
		}
		full[lake] = true
		pump[lake] = i
		result = append(result, -1)
	}
	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 && result[i] == -1 {
			result[i] = 1
		}
	}
	return result
}
