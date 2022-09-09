package main

func minOperations(logs []string) int {
	var depth int
	const moveUp = "../"
	const remain = "./"
	for _, log := range logs {
		if log == moveUp {
			if depth > 0 {
				depth--
			}
		} else if log != remain {
			depth++
		}
	}
	return depth
}
