package main

import (
	"fmt"

	"github.com/vincentvtran/goplayground/objects"
	"github.com/vincentvtran/goplayground/structures"
	"github.com/vincentvtran/goplayground/thread"
)

func main() {

	// **** Demo Threads ****
	fmt.Println("Running channel demo...")
	thread.DemoChannel()
	fmt.Println("Running mutex demo...")
	thread.DemoMutex()
	fmt.Println("Running semaphore demo...")
	thread.DemoSemaphore()

	// **** Demo Data Structures ****
	structures.DemoArrays()
	structures.DemoMaps()

	// **** Demo Objects ****
	objects.DemoCar()
}
