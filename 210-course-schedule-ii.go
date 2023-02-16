package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjacency := make([][]int, numCourses)
	preCourses := make([]int, numCourses)
	for _, line := range prerequisites {
		preCourses[line[0]]++
		adjacency[line[1]] = append(adjacency[line[1]], line[0])
	}

	var order []int
	for i := 0; i < numCourses; i++ {
		if preCourses[i] == 0 {
			order = append(order, i)
		}
	}
	for i := 0; i < len(order); i++ {
		for _, next := range adjacency[order[i]] {
			preCourses[next]--
			if preCourses[next] == 0 {
				order = append(order, next)
			}
		}
	}

	if len(order) != numCourses {
		return []int{}
	}
	return order
}
