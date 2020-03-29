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

func generatePath(index, step int, distance []int, linked [][]int) [][]int {
	if step == 2 {
		return [][]int{{index}}
	}
	var result [][]int
	for i := 0; i < len(linked[index]); i++ {
		j := linked[index][i]
		if distance[j] != step-1 {
			continue
		}
		paths := generatePath(j, step-1, distance, linked)
		for _, p := range paths {
			result = append(result, append([]int{index}, p...))
		}
	}
	return result
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var result [][]string
	endIndex := -1
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			endIndex = i
			break
		}
	}
	if endIndex == -1 {
		return result
	}

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

	var shortest int
	for i := 0; i < len(wordList); i++ {
		if distance[i] == 0 {
			continue
		}
		if diffCount(beginWord, wordList[i]) == 1 {
			if shortest == 0 || shortest > distance[i] {
				shortest = distance[i]
			}
		}
	}
	if shortest == 0 {
		return result
	} else if shortest == 1 {
		return [][]string{{beginWord, endWord}}
	}
	for i := 0; i < len(wordList); i++ {
		if distance[i] == shortest && diffCount(beginWord, wordList[i]) == 1 {
			r := generatePath(i, shortest, distance, linked)
			for _, indexes := range r {
				path := []string{beginWord}
				for _, j := range indexes {
					path = append(path, wordList[j])
				}
				path = append(path, endWord)
				result = append(result, path)
			}
		}
	}
	return result
}
