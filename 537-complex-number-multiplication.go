package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	n1 := strings.SplitN(num1, "+", 2)
	n2 := strings.SplitN(num2, "+", 2)
	a, _ := strconv.Atoi(n1[0])
	b, _ := strconv.Atoi(n1[1][:len(n1[1])-1])
	c, _ := strconv.Atoi(n2[0])
	d, _ := strconv.Atoi(n2[1][:len(n2[1])-1])
	real := a*c - b*d
	imaginary := b*c + a*d
	return fmt.Sprintf("%d+%di", real, imaginary)
}
