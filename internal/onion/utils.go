package onion

import "fmt"

func panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func warnf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
