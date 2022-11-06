package main

import "strings"

func interpret(command string) string {
	return strings.ReplaceAll(strings.ReplaceAll(command, "(al)", "al"), "()", "o")
}
