package main

import (
	"strconv"
	"time"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	getDays := func(date string) int {
		a, _ := strconv.Atoi(date[:2])
		b, _ := strconv.Atoi(date[3:])
		d := time.Date(2001, time.Month(a), b, 0, 0, 0, 0, time.UTC).Sub(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))
		return int(d.Hours()) / 24
	}
	getList := func(arrive, leave string) (r []int) {
		for start, end := getDays(arrive), getDays(leave); start <= end; start++ {
			r = append(r, start)
		}
		return
	}
	alice := make(map[int]bool)
	for _, d := range getList(arriveAlice, leaveAlice) {
		alice[d] = true
	}
	var result int
	for _, d := range getList(arriveBob, leaveBob) {
		if alice[d] {
			result++
		}
	}
	return result
}
