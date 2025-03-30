package onion_log

import "time"

type Timer struct {
	beginTime time.Time
}

func NewTimer() Timer {
	return Timer{
		beginTime: time.Now(),
	}
}

func (timer *Timer) Reset() {
	timer.beginTime = time.Now()
}

func (timer *Timer) Now() int64 {
	return time.Since(timer.beginTime).Microseconds()
}
