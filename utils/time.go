package utils

import (
	"time"
)

func ConvertTimeToString(value time.Time) string {
	return value.Format("2006-01-02 15:04:05")
}

func FillTheTime() time.Time {
	return time.Now()
}
