package repo

import (
	"fmt"
	"strconv"
	"time"
)

func toFloat(src interface{}) (float64, error) {
	fs, ok := src.(string)
	if !ok {
		return 0, fmt.Errorf("convert interface to string: %v", src)
	}
	return strconv.ParseFloat(fs, 64)
}

func toInt(src interface{}) (int, error) {
	is, ok := src.(string)
	if !ok {
		return 0, fmt.Errorf("convert interface to string: %v", src)
	}
	return strconv.Atoi(is)
}

func toTime(src interface{}) (time.Time, error) {
	ts, ok := src.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("convert interface to string: %v", src)
	}
	return time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
}

func toBool(src interface{}) (bool, error) {
	i, err := toInt(src)
	if err != nil {
		return false, err
	}
	ret := false
	if i != 0 {
		ret = true
	}
	return ret, nil
}
