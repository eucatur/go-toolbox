package log

import (
	"fmt"
	"time"
)

// TimeTrack — track the time spent by a function execution (defer recommended)
func TimeTrack(start time.Time, name string, log bool) {
	elapsed := time.Since(start)
	msg := fmt.Sprintf("%s took %s", name, elapsed)

	if log {
		File(time.Now().Format("performance/2006/01/02/15h.log"), msg)
	} else {
		fmt.Println(msg)
	}
}
