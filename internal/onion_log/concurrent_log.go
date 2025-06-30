package onion_log

import (
	"sort"
	"sync/atomic"
)

type ConcurrentLog struct {
	logs     []Log
	logIndex int32
	maxSize  int
}

func NewConcurrentLog(size int, timer *Timer) ConcurrentLog {
	logs := make([]Log, 0)
	for range size {
		logs = append(logs, NewLog(timer))
	}

	return ConcurrentLog{
		logs:     logs,
		logIndex: 0,
		maxSize:  size,
	}
}

func (cl *ConcurrentLog) Add(from, to uint32, comments ...any) {
	start := int(atomic.AddInt32(&cl.logIndex, 1))

	for i := range cl.maxSize {
		index := (start + i) % cl.maxSize
		log := &cl.logs[index]
		if log.TryAdd(from, to, comments...) {
			return
		}
	}
	cl.logs[start%cl.maxSize].Add(from, to, comments...)
}

func (cl *ConcurrentLog) MergeAndClear() []LogEntry {
	var allEntries []LogEntry

	for i := 0; i < cl.maxSize; i++ {
		entries := cl.logs[i].ExportAndClear()
		allEntries = append(allEntries, entries...)
	}

	sort.Slice(allEntries, func(i, j int) bool {
		return allEntries[i].Time < allEntries[j].Time
	})

	return allEntries
}
