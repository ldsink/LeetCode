package main

import (
	"strconv"
	"strings"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func getCommaIndex(s string) int {
	var count int
	for i, r := range s {
		if r == '[' {
			count++
		} else if r == ']' {
			count--
		} else if count == 0 && r == ',' {
			return i
		}
	}
	return -1
}

func deserialize(s string) *NestedInteger {
	n := NestedInteger{}
	if strings.HasPrefix(s, "[") {
		s = s[1 : len(s)-1]
		for idx := getCommaIndex(s); idx != -1; idx = getCommaIndex(s) {
			elem := deserialize(s[:idx])
			n.Add(*elem)
			s = s[idx+1:]
		}
		if len(s) > 0 {
			elem := deserialize(s)
			n.Add(*elem)
		}
	} else {
		i, _ := strconv.Atoi(s)
		n.SetInteger(i)
	}
	return &n
}
