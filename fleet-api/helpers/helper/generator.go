package helper

import "time"

func TimeGenerator() string {
	timeNow := time.Now().Format(time.RFC3339)
	return timeNow
}
