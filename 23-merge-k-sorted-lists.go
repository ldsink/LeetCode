package main

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
	for n := len(listNodeHeap) - 1; n > 0; {
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
	if (*result).Next != nil {
		addNode((*result).Next)
	}
	for current, node := result, popNode(); node != nil; node = popNode() {
		(*current).Next = node
		current = node
		if (*current).Next != nil {
			addNode((*current).Next)
		}
	}
	return result
}
