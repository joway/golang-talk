package utils

import (
	"fmt"
	"time"
)

func Log(msg string) (n int, err error) {
	return fmt.Println(msg)
}

func Logf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func Now() int64 {
	return time.Now().UnixNano() / 1000000
}
