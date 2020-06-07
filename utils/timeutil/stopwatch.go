package timeutil

import "time"

// Stopwatch 用于计算一段代码所运行的时间
type Stopwatch struct {
	start time.Time
	end   time.Time
}

// NewStopwatch 根据系统当前时间初始化 start 时间
func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		start: time.Now(),
	}
}

// ElapsedTime 返回自 Stopwatch 创建以来所经过的时间
func (sw *Stopwatch) ElapsedTime() time.Duration {
	sw.end = time.Now()
	latency := sw.end.Sub(sw.start)

	return latency
}
