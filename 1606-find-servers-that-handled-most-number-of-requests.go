package main

import (
	"container/heap"
)

type ServerStatus struct {
	id    int // 服务器编号
	ready int // 服务器可以提供服务的时间
}
type ServerHeap []ServerStatus

func (h ServerHeap) Len() int           { return len(h) }
func (h ServerHeap) Less(i, j int) bool { return h[i].ready < h[j].ready }
func (h ServerHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ServerHeap) Push(x interface{}) {
	*h = append(*h, x.(ServerStatus))
}
func (h *ServerHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func getDistance(k, i, j int) int {
	if j >= i%k {
		return j - (i % k)
	}
	return j + k - (i % k)
}

func busiestServers(k int, arrival []int, load []int) []int {
	var result []int

	// 任务数量不超过服务器数量，每个服务器都是最繁忙的
	if len(arrival) <= k {
		for i := 0; i < len(arrival); i++ {
			result = append(result, i)
		}
		return result
	}

	count := make([]int, k) // 服务器提供服务的次数
	ready := make([]int, k) // 服务器下次提供服务的时间
	serverHeap := ServerHeap{}
	for i := 0; i < k; i++ {
		count[i]++
		ready[i] = load[i]
		heap.Push(&serverHeap, ServerStatus{id: i, ready: arrival[i] + load[i]})
	}
	var spareServer []int // 空闲的服务器
	for i := k; i < len(arrival); i++ {
		// 没有服务器有空闲，跳过
		if len(spareServer) == 0 && arrival[i] < serverHeap[0].ready {
			continue
		}
		// 取出全部有空闲的服务器
		for serverHeap.Len() > 0 && arrival[i] >= serverHeap[0].ready {
			s := heap.Pop(&serverHeap)
			spareServer = append(spareServer, s.(ServerStatus).id)
		}
		// 找到最近的服务器
		idx := 0
		serverDistance := getDistance(k, i, spareServer[idx])
		for j := 1; j < len(spareServer); j++ {
			if d := getDistance(k, i, spareServer[j]); d < serverDistance {
				idx = j
				serverDistance = d
			}
		}
		// 添加这个服务器到堆中
		heap.Push(&serverHeap, ServerStatus{id: spareServer[idx], ready: arrival[i] + load[i]})
		count[spareServer[idx]]++
		// 从空闲服务器列表中移除
		if idx == 0 {
			spareServer = spareServer[1:]
		} else if idx == len(spareServer)-1 {
			spareServer = spareServer[:len(spareServer)-1]
		} else {
			spareServer = append(spareServer[:idx], spareServer[idx+1:]...)
		}
	}

	var max int
	for i := 0; i < k; i++ {
		if count[i] > max {
			max = count[i]
			result = []int{i}
		} else if count[i] == max {
			result = append(result, i)
		}
	}
	return result
}
