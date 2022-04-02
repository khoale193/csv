package b

import (
	"net/url"
	"strconv"
	"time"
)

// BindTime binding data from direct key on URL & convert to time.Time
func BindTime(values url.Values, key string, location *time.Location) time.Time {
	if value, ok := values[key]; ok && value[0] != "" {
		date, _ := strconv.ParseInt(value[0], 10, 64)
		return time.Unix(date, 0).In(location)
	}
	return time.Time{}
}

// BindString binding data from direct key on URL & convert to string
func BindString(values url.Values, key string) string {
	if value, ok := values[key]; ok && value[0] != "" {
		return value[0]
	}
	return ""
}

// BindSliceString binding data from direct key on URL & convert to slice of string
func BindSliceString(values url.Values, key string) []string {
	data := make([]string, 0)
	if value, ok := values[key]; ok {
		for _, bar := range value {
			data = append(data, bar)
		}
	}
	return data
}

// BindInt64 binding data from direct key on URL & convert to int64
func BindInt64(values url.Values, key string) int64 {
	if value, ok := values[key]; ok && value[0] != "" {
		if resultInt64, err := strconv.ParseInt(value[0], 10, 64); err == nil {
			return resultInt64
		}
	}
	return 0
}

func BindInt(values url.Values, key string) int {
	if value, ok := values[key]; ok && value[0] != "" {
		if resultInt64, err := strconv.ParseInt(value[0], 10, 64); err == nil {
			return int(resultInt64)
		}
	}
	return 0
}

// BindSliceInt binding data from direct key on URL & convert to slice of int
func BindSliceInt(values url.Values, key string) []int {
	data := make([]int, 0)
	if value, ok := values[key]; ok {
		for _, bar := range value {
			if id, err := strconv.Atoi(bar); err == nil {
				data = append(data, id)
			}
		}
	}
	return data
}
