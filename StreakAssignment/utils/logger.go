package utils

import (
	"log"
	"runtime"
)

// Loging for perfect tracibilty to see function name and line number
func Log(format string, args ...interface{}) {
	//getting stack trace
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Printf("[%s:%d] ", file, line) 
	}
	log.Printf(format, args...)
}
