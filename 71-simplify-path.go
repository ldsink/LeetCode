package main

import (
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var newPaths []string
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}
		if path == "." {
			continue
		}
		if path == ".." {
			if len(newPaths) > 0 {
				newPaths = newPaths[:len(newPaths)-1]
			}
			continue
		}
		newPaths = append(newPaths, path)
	}
	path = strings.Join(newPaths, "/")
	return "/" + path
}
