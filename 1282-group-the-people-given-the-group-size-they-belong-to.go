package main

func groupThePeople(groupSizes []int) [][]int {
	peopleGroup := make(map[int][]int)
	for people, size := range groupSizes {
		peopleGroup[size] = append(peopleGroup[size], people)
	}
	var result [][]int
	for size, peoples := range peopleGroup {
		for i := 0; i < len(peoples); i += size {
			result = append(result, peoples[i:i+size])
		}
	}
	return result
}
