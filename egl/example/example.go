package main

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	ctx := Create("Example", 640, 480)
	ctx.MainLoop()
}
