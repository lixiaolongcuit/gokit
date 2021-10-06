package timex

import "time"

//获取当前时间的秒级别时间戳(10位)
func TimestampSec() int64 {
	return time.Now().Unix()
}

//获取当前时间的毫秒级别时间戳(13位)
func TimestampMsec() int64 {
	return time.Now().UnixNano() / 1e6
}

//获取当前时间的纳秒级别时间戳(19位)
func TimestampNsec() int64 {
	return time.Now().UnixNano()
}

//秒类型(10位)时间戳转时间
func TimestampSec2Time(sec int64) time.Time {
	return time.Unix(sec, 0)
}

//毫秒类型(13位)时间戳转时间
func TimestampMsec2Time(msec int64) time.Time {
	return time.Unix(msec/1000, msec%1000*1e6)
}

//纳秒类型(19位)时间戳转时间
func TimestampNsec2Time(nsec int64) time.Time {
	return time.Unix(nsec/1e9, nsec%1e9)
}

//秒类型(10位)时间戳格式化成字符串
func TimestampSecFormat(sec int64, layout string) string {
	return TimestampSec2Time(sec).Format(layout)
}

//毫秒类型(13位)时间戳格式化成字符串
func TimestampMsecFormat(msec int64, layout string) string {
	return TimestampMsec2Time(msec).Format(layout)
}

//纳秒类型(19位)时间戳格式化成字符串
func TimestampNsecFormat(nsec int64, layout string) string {
	return TimestampNsec2Time(nsec).Format(layout)
}
