package main

func detectCycle(head *ListNode) *ListNode {
	for fast, slow, count := head, head, 0; fast != nil && fast.Next != nil && slow != nil; fast, slow = fast.Next.Next, slow.Next {
		if fast == slow {
			count++
			if count == 2 {
				for fast = head; fast != slow; fast, slow = fast.Next, slow.Next {
				}
				return fast
			}
		}
	}
	return nil
}
