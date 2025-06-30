package onion_log

import (
	"fmt"

	lock "github.com/viney-shih/go-lock"
)

type Log struct {
	entries []LogEntry
	lock    *lock.CASMutex
	timer   *Timer
}
type LogEntry struct {
	From     uint32
	To       uint32
	Comments []any
	Time     int64
}

func NewLog(timer *Timer) Log {
	return Log{
		entries: make([]LogEntry, 0),
		lock:    lock.NewCASMutex(),
		timer:   timer,
	}
}

func (log *Log) TryAdd(from, to uint32, comments ...any) bool {
	if log.lock.TryLock() {
		defer log.lock.Unlock()
		fmt.Printf("[%d][%d] %d\n", from, to, log.timer.Now())
		log.entries = append(log.entries, LogEntry{from, to, comments, log.timer.Now()})
		return true
	} else {
		return false
	}
}

func (log *Log) Add(from, to uint32, comments ...any) {
	log.lock.Lock()
	defer log.lock.Unlock()

	logEntry := LogEntry{from, to, comments, log.timer.Now()}
	log.entries = append(log.entries, logEntry)
	fmt.Println(logEntry)
}

func (log *Log) ExportAndClear() []LogEntry {
	log.lock.Lock()
	defer log.lock.Unlock()

	prevEntries := log.entries
	log.entries = make([]LogEntry, 0)
	return prevEntries
}

func (logEntry *LogEntry) String() string {
	comments := fmt.Sprint(logEntry.Comments...)
	return fmt.Sprintf("[%d][%d] %d %s", logEntry.From, logEntry.To, logEntry.Time, comments)
}
