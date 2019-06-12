package datetime

import "time"

const (
	DateTimeFormatLayout = "2006-01-02 15:04:05"
)

func Date(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
