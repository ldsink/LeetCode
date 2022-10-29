package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count int
	for _, item := range items {
		if ruleKey == "type" && item[0] == ruleValue {
			count++
		} else if ruleKey == "color" && item[1] == ruleValue {
			count++
		} else if ruleKey == "name" && item[2] == ruleValue {
			count++
		}
	}
	return count
}
