package utils

import (
	"fmt"
	"time"
)

func GetDateTime(d, t string) (time.Time, error) {
	dts := fmt.Sprintf("%s %s", d, t)
	if len(t) == 0 {
		dts = fmt.Sprintf("%s 00:00:00", d)
	}
	return time.ParseInLocation("2006-01-02 15:04:05", dts, time.Local)
}
