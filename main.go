package main

import (
	"fmt"

	"github.com/vincentvtran/goplayground/thread"
)

func main() {
	fmt.Println("Hello, Go.")
	thread.DemoChannel()
	thread.DemoSignals()
}
