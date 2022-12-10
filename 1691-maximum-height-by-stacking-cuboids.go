package main

import "sort"

func maxHeight(cuboids [][]int) int {
	const bitLen = 64
	type Item struct {
		size [3]int // 宽长高
		idx  [2]int // 编号
	}
	// 每个正方体基于底边，可以扩展出三种选项
	var items []Item
	for i, j := range cuboids {
		items = append(items, Item{size: [3]int{j[0], j[1], j[2]}, idx: [2]int{i}})
		items = append(items, Item{size: [3]int{j[0], j[2], j[1]}, idx: [2]int{i}})
		items = append(items, Item{size: [3]int{j[1], j[2], j[0]}, idx: [2]int{i}})
	}
	// 统一处理一下选项，方便后面计算
	for i, j := range items {
		// 确保宽 <= 长
		if j.size[0] > j.size[1] {
			items[i].size[0], items[i].size[1] = items[i].size[1], items[i].size[0]
		}
		// idx 从数值转换成二进制位，方便后面比较
		idx := items[i].idx[0]
		items[i].idx[0] = 0
		items[i].idx[(idx / bitLen)] = 1 << (idx % bitLen)
	}
	// 从大到小排序
	sort.Slice(items, func(i, j int) bool {
		for k := 2; k >= 0; k-- {
			if items[i].size[k] > items[j].size[k] {
				return true
			} else if items[i].size[k] < items[j].size[k] {
				return false
			}
		}
		return false
	})
	// 开始搜索
	type Condition struct {
		height int
		items  [2]int
	}
	conditions := make([]Condition, len(items))
	// j 是否可以堆叠在 i 上，需要满足的条件为 j 里面不含 i，且 j 的长宽高都小于 i
	canPlace := func(i, j int) bool {
		// j 的最佳堆叠已经包含了 i
		for k := 0; k < 2; k++ {
			if conditions[j].items[k]&items[i].idx[k] != 0 {
				return false
			}
		}
		for k := 0; k < 3; k++ {
			if items[j].size[k] > items[i].size[k] {
				return false
			}
		}
		return true
	}
	for i := len(items) - 1; i >= 0; i-- {
		conditions[i].height = items[i].size[2]
		conditions[i].items = items[i].idx
		for j := i + 1; j < len(items); j++ {
			if conditions[i].height < items[i].size[2]+conditions[j].height && canPlace(i, j) {
				conditions[i].height = items[i].size[2] + conditions[j].height
				for k := 0; k < 2; k++ {
					conditions[i].items[k] = conditions[j].items[k] ^ items[i].idx[k]
				}
			}
		}
	}
	var result int
	for i := 0; i < len(conditions); i++ {
		if result < conditions[i].height {
			result = conditions[i].height
		}
	}
	return result
}
