package main

import (
	"sort"
)

type result332 struct {
	order []string
}

func dfs332(graph map[string][]string, airport string, result *result332) {
	for ; len(graph[airport]) > 0; {
		port := graph[airport][0]
		graph[airport] = graph[airport][1:]
		dfs332(graph, port, result)
	}
	result.order = append(result.order, airport)
}

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}
	for airport, _ := range graph {
		sort.Strings(graph[airport])
	}
	var result result332
	dfs332(graph, "JFK", &result)
	for i, j := 0, len(result.order)-1; i < j; i, j = i+1, j-1 {
		result.order[i], result.order[j] = result.order[j], result.order[i]
	}
	return result.order
}
