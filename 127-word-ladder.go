package main

func diffCount(s1, s2 string) int {
	var count int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	endIndex := -1
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			endIndex = i
			break
		}
	}
	if endIndex == -1 {
		return 0
	}

	wordList = append([]string{beginWord}, wordList...)
	endIndex++
	var linked [][]int
	for i := 0; i < len(wordList); i++ {
		linked = append(linked, []int{})
	}
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if diffCount(wordList[i], wordList[j]) == 1 {
				linked[i] = append(linked[i], j)
				linked[j] = append(linked[j], i)
			}
		}
	}

	distance := make([]int, len(wordList))
	distance[endIndex] = 1
	for stack := []int{endIndex}; len(stack) > 0; stack = stack[1:] {
		index := stack[0]
		for i := 0; i < len(linked[index]); i++ {
			j := linked[index][i]
			if distance[j] > 0 {
				continue
			}
			distance[j] = distance[index] + 1
			stack = append(stack, j)
		}
	}
	return distance[0]
}
