package logger

import (
	"fmt"
	"runtime"
)

func PrintStack() {
	var buf [2 << 10]byte
	fmt.Println(string(buf[:runtime.Stack(buf[:], true)]))
}
