package main

import (
	"fmt"

	"github.com/vincentvtran/goplayground/thread"
)

func main() {
	fmt.Println("Running channel demo...")
	thread.DemoChannel()
	fmt.Println("Running mutex demo...")
	thread.DemoMutex()
	fmt.Println("Running semaphore demo...")
	thread.DemoSemaphore()
}
