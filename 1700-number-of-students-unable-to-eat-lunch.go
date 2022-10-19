package main

func countStudents(students []int, sandwiches []int) int {
	var circular, square int
	for _, student := range students {
		if student == 0 {
			circular++
		} else {
			square++
		}
	}
	var i int
	for ; i < len(sandwiches); i++ {
		if sandwiches[i] == 0 && circular > 0 {
			circular--
		} else if sandwiches[i] == 1 && square > 0 {
			square--
		} else {
			break
		}
	}
	return len(sandwiches) - i
}
