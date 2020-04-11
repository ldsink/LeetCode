package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	netValue := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		netValue[i] = gas[i] - cost[i]
	}

	for station := 0; station < len(gas); station++ {
		if netValue[station] < 0 || gas[station]-cost[station] < 0 {
			continue
		}
		gasTank := netValue[station]
		netValue[station] = 0
		for i := station + 1; ; i++ {
			if i == len(gas) {
				i = 0
			}
			if i == station {
				return station
			}
			if r := gasTank + netValue[i]; r >= 0 {
				netValue[i] = 0
				gasTank = r
			} else {
				netValue[i] = r
				break
			}
		}
	}
	return -1
}

func main() {
	var gas, cost []int
	gas = []int{2}
	cost = []int{2}
	fmt.Println(canCompleteCircuit(gas, cost))
	gas = []int{1, 2, 3, 4, 5}
	cost = []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
