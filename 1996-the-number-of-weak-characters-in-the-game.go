package main

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] < properties[j][1])
	})
	attack1 := properties[len(properties)-1][0]
	defense1 := properties[len(properties)-1][1]
	var attack2, defense2, count int
	for i := len(properties) - 2; i >= 0; i-- { // 按照攻击力从大到小逐个检查
		a, d := properties[i][0], properties[i][1]
		// 首先看当前的点是不是弱者，如果是，那直接往前走
		if (a < attack1 && d < defense1) || (attack2 != 0 && a < attack2 && d < defense2) {
			count++
			continue
		}
		// 其次看能不能更新最大防御。需要保留上一次的最大防御。因为上一次的攻击可能比之后的大。
		if d > defense1 {
			attack2, defense2 = attack1, defense1
			attack1, defense1 = a, d
		}
	}
	return count
}
