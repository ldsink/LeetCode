package main

import "container/heap"

type ClassStatus struct {
	pass, total int
	weight      float64
}
type MaxHeap []ClassStatus

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i].weight > (*h)[j].weight }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(ClassStatus))
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	// 只有 1 个班级，提前处理退出，保证后面的堆至少有 2 个班级
	if len(classes) == 1 {
		return float64(classes[0][0]+extraStudents) / float64(classes[0][1]+extraStudents)
	}
	getWeight := func(pass, total int) float64 {
		return float64(pass+1)/float64(total+1) - float64(pass)/float64(total)
	}
	h := &MaxHeap{}
	for _, class := range classes {
		heap.Push(h, ClassStatus{pass: class[0], total: class[1], weight: getWeight(class[0], class[1])})
	}
	for extraStudents > 0 {
		class := heap.Pop(h).(ClassStatus) // 先取出添加学生效果最好的班级
		var student int
		for ; class.weight >= (*h)[0].weight && student < extraStudents; student += 1 {
			class.pass += 1
			class.total += 1
			class.weight = getWeight(class.pass, class.total)
		}
		extraStudents -= student
		heap.Push(h, class)
	}
	var result float64
	for i := 0; i < len(*h); i++ {
		result += float64((*h)[i].pass) / float64((*h)[i].total)
	}
	return result / float64(len(*h))
}
