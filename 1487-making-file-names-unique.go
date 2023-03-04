package main

import "fmt"

func getFolderNames(names []string) []string {
	exists := make(map[string]int)
	var result []string
	var newName string
	for _, name := range names {
		if idx, ok := exists[name]; !ok {
			result = append(result, name)
			exists[name] = 1
		} else {
			for ; ; idx++ {
				newName = fmt.Sprintf("%s(%d)", name, idx)
				if _, ok := exists[newName]; !ok {
					result = append(result, newName)
					exists[newName] = 1
					exists[name] = idx + 1
					break
				}
			}
		}
	}
	return result
}
