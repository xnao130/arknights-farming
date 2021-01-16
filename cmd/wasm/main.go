package main

import (
	"syscall/js"
)

func gofunc(this js.Value, args []js.Value) interface{} {
	return 100
}

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("gofunc", js.FuncOf(gofunc))
	<-done
}
