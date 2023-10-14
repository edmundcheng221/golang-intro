package time

import (
	"fmt"
	"time"
)

var print = fmt.Println

func GetTime() time.Time {
	return time.Now()
}

func GetEpoch() {
	now := time.Now()
	print(now)
	print(now.Unix())      // epoch time in seconds
	print(now.UnixMilli()) // epoch time in milliseconds
	print(now.UnixNano())  // epoch time in nanoseconds
}

func FormatTime() {
	t := time.Now()
	print(t.Format("3:04PM")) // reference formats to show pattern
	print(t.Format("Mon Jan _2 15:04:05 2006"))
	print(t.Format("2006-01-02T15:04:05.999999-07:00"))
}
