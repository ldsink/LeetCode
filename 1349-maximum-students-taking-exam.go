package main

import (
	"sort"
)

const DISABLE = byte('#')
const PENDING = byte('.')
const USED = byte('1')

func getAllCases(line []byte, prefix []byte, index int) []string {
	for i := index; i < len(line); i++ {
		if line[i] == DISABLE {
			prefix = append(prefix, DISABLE)
		} else {
			// 不使用情况下的长度
			result := getAllCases(line, append(prefix, PENDING), i+1)
			// 如果当前位置可以使用的话
			if i == 0 || prefix[i-1] != USED {
				result = append(result, getAllCases(line, append(prefix, USED), i+1)...)
			}
			return result
		}
	}
	return []string{string(prefix)}
}

func canNotCheating(base, front string) bool {
	for i := 0; i < len(base); i++ {
		if base[i] == USED {
			if i != 0 && front[i-1] == USED {
				return false
			}
			if i+1 < len(base) && front[i+1] == USED {
				return false
			}
		}
	}
	return true
}

func maxStudents(seats [][]byte) int {
	type seatsCount struct {
		condition string
		count     int
	}
	var prev, curr []seatsCount
	for i := 0; i < len(seats); i++ {
		// 拿到当前层级所有合法的情况
		cases := getAllCases(seats[i], []byte{}, 0)

		// 设置当前层级的学生数
		curr = []seatsCount{}
		for _, condition := range cases {
			count := 0
			for _, r := range condition {
				if byte(r) == USED {
					count++
				}
			}
			curr = append(curr, seatsCount{condition: condition, count: count})
		}

		// 从第二层起，需要和上一层做检查，更新允许的最大值
		if i > 0 {
			for idx, base := range curr {
				// 从后往前匹配，因为 prev 为顺序，所以只要成功就可以停止
				for j := len(prev) - 1; j >= 0; j-- {
					if canNotCheating(base.condition, prev[j].condition) {
						curr[idx].count += prev[j].count
						break
					}
				}
			}
		}

		// 当前层变成下一次的上一层
		prev = curr

		// 排序，这样可以减少比较次数
		sort.Slice(prev, func(i, j int) bool {
			return prev[i].count < prev[j].count
		})
	}
	return prev[len(prev)-1].count
}
