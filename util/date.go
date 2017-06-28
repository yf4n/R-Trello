package util

import (
	"log"
	"time"
)

func GetThisWeekDate() string {
	p := log.Println
	log.Println("This is a date range")
	now := time.Now()

	p(time.Now())
	p(now.Weekday())

	return ""
}
