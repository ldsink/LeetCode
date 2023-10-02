package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	type Transaction struct {
		name, city   string
		amount, time int
	}
	var data []Transaction
	for _, s := range transactions {
		parts := strings.Split(s, ",")
		time, _ := strconv.Atoi(parts[1])
		amount, _ := strconv.Atoi(parts[2])
		data = append(data, Transaction{name: parts[0], time: time, amount: amount, city: parts[3]})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].name < data[j].name || (data[i].name == data[j].name && data[i].time < data[j].time)
	})
	invalid := make(map[int]bool)
	for i, t := range data {
		if t.amount > 1000 {
			invalid[i] = true
		}
		for j := i + 1; j < len(data); j++ {
			if data[j].name != data[i].name {
				break
			}
			if data[i].city != data[j].city && data[j].time-data[i].time <= 60 {
				invalid[i], invalid[j] = true, true
			}
		}
	}
	var result []string
	for i := range invalid {
		result = append(result, fmt.Sprintf("%s,%d,%d,%s", data[i].name, data[i].time, data[i].amount, data[i].city))
	}
	return result
}
