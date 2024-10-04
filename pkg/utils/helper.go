package utils

import "time"

func ConvertStrToTime(publishedAt string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, publishedAt)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
