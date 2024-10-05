package utils

import "time"

// Converts a string representation of time in RFC3339 format to a time.Time object. (the google api returns a string)
func ConvertStrToTime(publishedAt string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, publishedAt)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
