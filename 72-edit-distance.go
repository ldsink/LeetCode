package main

func _dynamicSearch(index1, index2 int, rune1, rune2 []rune, searched *map[int]map[int]int) int {
	distances := *searched
	if _, ok := distances[index1]; ok {
		if distance, ok := distances[index1][index2]; ok {
			return distance
		}
	} else {
		distances[index1] = make(map[int]int)
	}

	if index1 == len(rune1) || index2 == len(rune2) {
		distances[index1][index2] = (len(rune1) - index1) + (len(rune2) - index2)
	} else if rune1[index1] == rune2[index2] {
		distances[index1][index2] = _dynamicSearch(index1+1, index2+1, rune1, rune2, searched)
	} else {
		distance := _dynamicSearch(index1, index2+1, rune1, rune2, searched) + 1
		if d := _dynamicSearch(index1+1, index2+1, rune1, rune2, searched) + 1; distance > d {
			distance = d
		}
		if d := _dynamicSearch(index1+1, index2, rune1, rune2, searched) + 1; distance > d {
			distance = d
		}
		distances[index1][index2] = distance
	}
	return distances[index1][index2]
}

func minDistance(word1 string, word2 string) int {
	searched := make(map[int]map[int]int)
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}
	rs1 := []rune(word1)
	rs2 := []rune(word2)
	return _dynamicSearch(0, 0, rs1, rs2, &searched)
}
