package main

import (
	"fmt"
	"sync"
	"time"
)

// This example simulates a common concurrency issue that can be tricky in languages like Rust,
// where explicit synchronization is often required and can lead to complex debugging. 
// Go's built-in goroutines and channels offer a simpler, often more readable approach to concurrency.

func main() {
	fmt.Println("Starting simulation...")

	// In Rust, a similar scenario might involve complex mutexes or other synchronization primitives
	// that could lead to deadlocks or race conditions if not managed perfectly. 
	// Here, we use channels for safe communication between goroutines.

	var wg sync.WaitGroup
	dataChannel := make(chan int)

	// Producer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("Producer sending: %d\n", i)
			dataChannel <- i // Send data to the channel
			time.Sleep(100 * time.Millisecond) // Simulate work
		}
		close(dataChannel) // Close the channel when done sending
	}()

	// Consumer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		// This loop will automatically exit when the channel is closed and empty.
		// This avoids the need for explicit checks or complex synchronization logic that could cause crashes.
		for data := range dataChannel {
			fmt.Printf("Consumer received: %d\n", data)
			time.Sleep(150 * time.Millisecond) // Simulate processing
		}
		fmt.Println("Consumer finished.")
	}()

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("Simulation finished.")
}
