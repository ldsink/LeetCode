package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s", t.Weekday())
}
