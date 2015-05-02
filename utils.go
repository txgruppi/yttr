package yttr

import (
	"net/url"
	"time"
)

var (
	dateFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
)

func formattedDate() string {
	now := time.Now().In(time.FixedZone("UTC", 0))
	return now.Format(dateFormat)
}

func parseDate(date string) (time.Time, error) {
	return time.Parse(dateFormat, date)
}

func encodeURIComponent(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}
