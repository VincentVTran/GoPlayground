package thread

import (
	"fmt"
	"sync"
	"time"
)

func DemoChannel() {
	mutexChannelDemo()
	mutexSemaphoreDemo()
}

func mutexChannelDemo() {
	ch := make(chan int, 1) // Create a channel of type int with a buffer size of 1 (optional, can be unbuffered)
	// Run the goroutine and pass the channel
	go uploadChannel(1, ch)

	// Blocks the main go routine until it receives a value from the channel (sleep)
	val := <-ch
	fmt.Println(val)
}

func mutexSemaphoreDemo() {
	ch := make(chan int, 3) // Create a channel of type int with a buffer size of 1 (optional, can be unbuffered)
	// Run the goroutine and pass the channel
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go uploadChannel(i, ch)
	}

	// Blocks the main go routine until it receives a value from the channel (sleep)
	// Run the goroutine and pass the channel
	for i := 0; i < 10; i++ {
		<-ch
		wg.Done()
	}
}

func uploadChannel(threadId int, desiredChannel chan int) {
	// Simulate workload
	time.Sleep(2 * time.Second)
	desiredChannel <- 10
	fmt.Println("Goroutine", threadId, "finished processing")
}
