package timeutil

import "time"

// time layout format
// 这里只定义 layout 格式，请按需选择即可，格式化输出请使用： t.Format(layout)
const (
	Layout2DateDay    = "2006-01-02"
	Layout2DateHour   = "2006-01-02 15"
	Layout2DateMinute = "2006-01-02 15:04"
	Layout2DateTime   = "2006-01-02 15:04:05"
	Layout2StampMilli = "2006-01-02 15:04:05.000"
	Layout2StampMicro = "2006-01-02 15:04:05.000000"
	Layout2StampNano  = "2006-01-02 15:04:05.000000000"
)

// Unix 时间戳（精确到秒）
func Unix(t time.Time) int64 {
	return t.Unix()
}

// UnixMilli 时间戳（精确到毫秒）
func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// UnixMicro 时间戳（精确到微秒）
func UnixMicro(t time.Time) int64 {
	return t.UnixNano() / 1e3
}

// UnixNano 时间戳（精确到纳秒）
func UnixNano(t time.Time) int64 {
	return t.UnixNano()
}

// DayRange 根据传入时间，获取当天的起始时间点
func DayRange(t time.Time, loc *time.Location) (time.Time, time.Time) {
	y, m, d := t.Date()
	s := time.Date(y, m, d, 0, 0, 0, 0, loc)
	e := s.AddDate(0, 0, 1).Add(-1 * time.Second)
	return s, e
}

// MonthRange 根据传入时间，获取当月的起始时间点
func MonthRange(t time.Time, loc *time.Location) (time.Time, time.Time) {
	y, m, _ := t.Date()
	s := time.Date(y, m, 1, 0, 0, 0, 0, loc)
	e := s.AddDate(0, 1, 0).Add(-1 * time.Second)
	return s, e
}

// PositiveUnix 获取 time.Time 的时间戳, 若time小于1970-01-01 00:00:00 则返回0
func PositiveUnix(t time.Time) int64 {
	var tm int64
	if t.Unix() > 0 {
		tm = t.Unix()
	}

	return tm
}

// RangeN 从指定时间开始，按 step 时间步长，生成 n 个 time.Time时间切片
// 返回 [start, start+n*step] 范围内的时间切片
func RangeN(start time.Time, step time.Duration, n int) []time.Time {
	if n < 1 {
		return nil
	}

	return Range(start, start.Add(time.Duration(n-1)*step), step)
}

// Range 根据 [start,end]指定的范围，按 step 间隔生成时间切片
func Range(start, end time.Time, step time.Duration) []time.Time {
	out := make([]time.Time, 0)
	for end.After(start) || end.Equal(start) {
		out = append(out, start)
		start = start.Add(step)
	}

	return out
}
