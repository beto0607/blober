package utils

import "time"

func UTCTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
