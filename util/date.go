package util

import (
	"log"
	"time"
)

var p = log.Println

var days = map[string]int{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

// GetDateString 获取当前时间格式化字符串
func GetDateString(ts int64) string {
	format := "2006/01/02 15:04:05"

	return time.Unix(ts, 0).Format(format)
}

// GetDateStringWithFormat 格式化当前时间戳
func GetDateStringWithFormat(ts int64, format string) string {
	return time.Unix(ts, 0).Format(format)
}

// GetTodayDateString 获得当天的时间字符串
func GetTodayDateString() string {
	ts := time.Now().Unix()

	return GetDateStringWithFormat(ts, "2006-01-02")
}

// GetWeekDateRange 返回本周开头和结尾的时间戳
func GetWeekDateRange(t time.Time) (int64, int64) {
	weekday := days[time.Weekday.String(t.Weekday())]
	ts := t.Unix()/86400*86400 - 28800

	endTs := ts + int64((7-weekday)%7*86400)
	startTs := endTs - int64(518400)
	endTs += 86399

	return startTs, endTs
}
