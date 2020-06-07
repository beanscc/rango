package timeutil

import (
	"testing"
	"time"
)

func Test_StopWatch_ElapsedTime(t *testing.T) {
	stopwatch := NewStopwatch()
	time.Sleep(500 * time.Millisecond)
	latency := stopwatch.ElapsedTime()
	t.Logf("latency: %v", latency)

	time.Sleep(2 * time.Second)
	latency = stopwatch.ElapsedTime()
	t.Logf("latency: %v", latency)
}
