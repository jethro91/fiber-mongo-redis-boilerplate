package util

import "time"

// TimeNowUnixEpoch using milisecond since unix epoch
func TimeNowUnixEpoch() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
