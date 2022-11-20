package main

import (
	"math"
	"sort"
)

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	linked := make([][]int, n)
	for _, road := range roads {
		linked[road[0]] = append(linked[road[0]], road[1])
		linked[road[1]] = append(linked[road[1]], road[0])
	}
	distance, nextCity := make([]int, n), make([]int, n) // 到首都的距离，首都的下一座城市
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt
	}
	distance[0] = 0
	for queue := []int{0}; len(queue) > 0; queue = queue[1:] {
		for _, city := range linked[queue[0]] {
			if distance[city] > distance[queue[0]]+1 {
				distance[city] = distance[queue[0]] + 1
				nextCity[city] = queue[0]
				queue = append(queue, city)
			}
		}
	}
	var cities []int // cities 内不包含首都
	for i := 1; i < n; i++ {
		cities = append(cities, i)
	}
	sort.Slice(cities, func(i, j int) bool {
		return distance[cities[i]] > distance[cities[j]]
	})
	var result int64
	hasCar, hasPeople := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		hasPeople[i] = 1
	}
	for _, city := range cities {
		// 车辆不足，需要增加车
		for ; hasCar[city]*seats < hasPeople[city]; hasCar[city]++ {
		}
		// 车辆过多，可以减少车
		for ; (hasCar[city]-1)*seats >= hasPeople[city]; hasCar[city]-- {
		}
		result += int64(hasCar[city]) // 车开到下一个城市的油
		// 下一个城市会有这个城市过去的车和代表
		hasCar[nextCity[city]] += hasCar[city]
		hasPeople[nextCity[city]] += hasPeople[city]
	}
	return result
}
