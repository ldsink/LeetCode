package main

import "fmt"

var listNodeHeap []*ListNode

func popNode() *ListNode {
	if len(listNodeHeap) == 0 {
		return nil
	} else if len(listNodeHeap) == 1 {
		ret := listNodeHeap[0]
		listNodeHeap = []*ListNode{}
		return ret
	}

	top := listNodeHeap[0]
	end := listNodeHeap[len(listNodeHeap)-1]
	listNodeHeap = listNodeHeap[0 : len(listNodeHeap)-1]
	listNodeHeap[0] = end
	for n := 0; n < len(listNodeHeap); {
		l := (n << 1) + 1
		r := (n << 1) + 2
		if l < len(listNodeHeap) && r < len(listNodeHeap) {
			var min int
			if (*(listNodeHeap[l])).Val > (*(listNodeHeap[r])).Val {
				min = r
			} else {
				min = l
			}
			if (*(listNodeHeap[n])).Val > (*(listNodeHeap[min])).Val {
				listNodeHeap[n], listNodeHeap[min] = listNodeHeap[min], listNodeHeap[n]
				n = min
				continue
			}
		} else if l < len(listNodeHeap) && (*(listNodeHeap[n])).Val > (*(listNodeHeap[l])).Val {
			listNodeHeap[n], listNodeHeap[l] = listNodeHeap[l], listNodeHeap[n]
			n = l
			continue
		} else if r < len(listNodeHeap) && (*(listNodeHeap[n])).Val > (*(listNodeHeap[r])).Val {
			listNodeHeap[n], listNodeHeap[r] = listNodeHeap[r], listNodeHeap[n]
			n = r
			continue
		}
		break
	}
	return top
}

func addNode(node *ListNode) {
	listNodeHeap = append(listNodeHeap, node)
	n := len(listNodeHeap) - 1
	for ; n > 0; {
		p := (n - 1) >> 1
		if (*(listNodeHeap[n])).Val < (*(listNodeHeap[p])).Val {
			listNodeHeap[n], listNodeHeap[p] = listNodeHeap[p], listNodeHeap[n]
			n = p
			continue
		}
		break
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	listNodeHeap = []*ListNode{}
	for _, node := range lists {
		if node != nil {
			addNode(node)
		}
	}
	var result *ListNode
	result = popNode()
	if result == nil {
		return result
	}
	current := result
	if (*current).Next != nil {
		addNode((*current).Next)
	}
	for ; len(listNodeHeap) != 0; {
		node := popNode()
		if (*node).Next != nil {
			addNode((*node).Next)
		}
		(*current).Next = node
		current = node
	}
	return result
}

func main() {
	testCase := [][]int{
		[]int{-8, -7, -7, -5, 1, 1, 3, 4},
		[]int{-2},
		[]int{-10, -10, -7, 0, 1, 3},
		[]int{2},
	}
	var lists []*ListNode
	for _, array := range testCase {
		r := constructListNode(array)
		lists = append(lists, r)
	}
	result := mergeKLists(lists)
	for ; result != nil; {
		fmt.Print((*result).Val, " ")
		result = (*result).Next
	}
}
