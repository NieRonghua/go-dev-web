package timeutils

import "time"

func CurrentDay() string {
	return time.Now().UTC().Add(offset).Format("20060102")
}

func FormatZap(t time.Time) string {
	return t.UTC().Add(offset).Format(time.RFC3339Nano)
}
