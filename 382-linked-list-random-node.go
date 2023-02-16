package main

import (
	"crypto/rand"
	"math/big"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	list []int
}

/*
  - @param head The linked list's head.
    Note that the head is guaranteed to be not null, so it contains at least one node.
*/
func Constructor(head *ListNode) Solution {
	solution := Solution{}
	for ; head != nil; head = head.Next {
		solution.list = append(solution.list, head.Val)
	}
	return solution
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(this.list))))
	return this.list[int(n.Int64())]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
