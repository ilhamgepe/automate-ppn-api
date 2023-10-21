package utils

import "time"

func CurrentTimeFormatted() string {
	currentTime := time.Now().Add(2 * time.Hour)
	formattedTime := currentTime.Format("2006-01-02T15:04:05.999999999-07:00")
	return formattedTime
}
