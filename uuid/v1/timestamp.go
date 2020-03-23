package v1

import (
	"time"
)

// The number of 100-nanoseconds from "1582-10-15 00:00:00 +0000 UTC" to "1970-01-01 00:00:00 +0000 UTC".
const unixToUUID = 122192928000000000

// uuidTimestamp returns the number of 100-nanoseconds elapsed since "1582-10-15 00:00:00 +0000 UTC"(UUID-epoch).
func uuidTimestamp() int64 {
	timeNow := time.Now()
	return timeNow.Unix()*1e7 + int64(timeNow.Nanosecond())/100 + unixToUUID
}

// tillNext100nano spin wait till next 100-nanosecond.
func tillNext100nano(lastTimestamp int64) int64 {
	timestamp := uuidTimestamp()
	for timestamp <= lastTimestamp {
		timestamp = uuidTimestamp()
	}
	return timestamp
}
