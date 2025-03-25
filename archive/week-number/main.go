package main

import (
	"fmt"
	"time"
)

func GetISOWeekNumber(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

func main() {
	currentTime := time.Now()
	weekNumber := GetISOWeekNumber(currentTime)
	fmt.Printf("Current Week Number: %d\n", weekNumber)
}
