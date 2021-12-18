package main

func canCross(stones []int) bool {
	conditions := make([]map[int]bool, len(stones))
	indexes := make(map[int]int)
	for i := 0; i < len(stones); i++ {
		conditions[i] = make(map[int]bool)
		indexes[stones[i]] = i
	}
	conditions[0][1] = true
	end := stones[len(stones)-1]
	for i := 0; i < len(conditions); i++ {
		for k, _ := range conditions[i] {
			// 直接跳到了终点
			if stones[i]+k == end {
				return true
			}
			// 下一个石头存在，则下一个石头上，跳跃距离增加新的可能性。
			if idx, ok := indexes[stones[i]+k]; ok {
				conditions[idx][k] = true
				conditions[idx][k+1] = true
				if k > 1 {
					conditions[idx][k-1] = true
				}
			}
		}
		// 释放内存
		conditions[i] = nil
	}
	return false
}
