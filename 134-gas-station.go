package main

func canCompleteCircuit(gas []int, cost []int) int {
	netValue := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		netValue[i] = gas[i] - cost[i]
	}

	for station := 0; station < len(gas); station++ {
		if (netValue[station] == 0 && gas[station] != cost[station]) || netValue[station] < 0 {
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
			if netValue[i] == 0 {
				continue
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
