package date_utils

import "time"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(time.RFC3339)
}
