package datautil

import (
	"time"
)

func GetRFC3339TimeString() string {
	return time.Now().Format(time.RFC3339)
}
