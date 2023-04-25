package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	type People struct {
		name   string
		height int
	}
	var peoples []People
	for i, name := range names {
		peoples = append(peoples, People{name: name, height: heights[i]})
	}
	sort.Slice(peoples, func(i, j int) bool {
		return peoples[i].height > peoples[j].height
	})
	var result []string
	for _, p := range peoples {
		result = append(result, p.name)
	}
	return result
}
