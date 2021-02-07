package handler

import "time"

func parseStringToTime(timeString string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, timeString)
	return t, err

}
