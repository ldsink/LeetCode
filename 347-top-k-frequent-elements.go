package main

func adjustHeap(idx int, heap []int, count map[int]int) {
	left := idx*2 + 1
	if left < len(heap) && count[heap[idx]] > count[heap[left]] {
		heap[idx], heap[left] = heap[left], heap[idx]
		adjustHeap(left, heap, count)
	}
	right := left + 1
	if right < len(heap) && count[heap[idx]] > count[heap[right]] {
		heap[idx], heap[right] = heap[right], heap[idx]
		adjustHeap(right, heap, count)
	}
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n] += 1
	}

	var heap []int
	for n, c := range count {
		if len(heap) < k {
			heap = append(heap, n)
			for i := len(heap) - 1; i != 0; i = (i - 1) / 2 {
				if count[heap[i]] >= count[heap[(i-1)/2]] {
					break
				}
				heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
			}
		} else if count[heap[0]] >= c {
			continue
		} else {
			heap[0] = n
			adjustHeap(0, heap, count)
		}
	}
	return heap
}
