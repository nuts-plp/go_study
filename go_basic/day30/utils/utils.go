package utils

import (
	"strconv"
	"time"
)

func GetUnix() string {
	now := time.Now().Unix()
	return strconv.FormatInt(now, 10)
}
func GetDay() string {
	now := time.Now()
	return now.Format("2006-01-02")
}
