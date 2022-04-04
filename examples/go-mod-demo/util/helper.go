package util

import "time"

// UnixTime time()
func UnixTime() int64 {
	return time.Now().Unix()
}

// 时间戳转成 datetime 字符串 YY-mm-dd
func UnixTime2Date(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02")
}

// 时间戳转成 datetime 字符串 YY-mm-dd HH:ii:ss
func UnixTime2DateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
