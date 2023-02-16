package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacency := make([][]int, numCourses)
	preCourses := make([]int, numCourses)
	for _, line := range prerequisites {
		preCourses[line[0]]++
		adjacency[line[1]] = append(adjacency[line[1]], line[0])
	}

	var taken []int
	for i := 0; i < numCourses; i++ {
		if preCourses[i] == 0 {
			taken = append(taken, i)
		}
	}
	for i := 0; i < len(taken); i++ {
		for _, next := range adjacency[taken[i]] {
			preCourses[next]--
			if preCourses[next] == 0 {
				taken = append(taken, next)
			}
		}
	}

	return len(taken) == numCourses
}
