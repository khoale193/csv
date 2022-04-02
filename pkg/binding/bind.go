package b

import (
	"time"

	"github.com/khoale193/csv/pkg/e"
)

func Int64toTime(bar int64) time.Time {
	if bar > 0 {
		timezone, _ := time.LoadLocation(e.DefaultTimezone)
		return time.Unix(bar, 0).In(timezone)
	}
	return time.Time{}
}
