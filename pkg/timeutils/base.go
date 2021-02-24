package timeutils

import "time"

var (
	offset time.Duration
)

func init() {
	offset, _ = time.ParseDuration("+08.00h")
}
