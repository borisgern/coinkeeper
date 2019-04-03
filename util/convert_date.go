package util

import (
	"time"
)

func ToUnixFormat(date string) (int64, error) {
	d, err := time.Parse("1/2/2006", date)
	return d.Unix(), err
}
