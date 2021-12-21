package main

import (
	"time"
)

func dayOfYear(date string) int {
	end, _ := time.Parse("2006-01-02", date)
	start, _ := time.Parse("2006-01-02", date[0:4]+"-01-01")
	delta := end.Sub(start)
	return int(delta.Hours())/24 + 1
}
