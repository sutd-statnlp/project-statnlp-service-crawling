package util

import "time"

// GetCurrentTimestamp .
func GetCurrentTimestamp() string {
	return time.Now().Format("20060102150405")
}
