package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a WaitGroup to wait for task to finish
	var wg sync.WaitGroup

	// Add the number of task to the WaitGroup
	wg.Add(2)

	// Launch the first task
	go func() {
		// Defer the Done method to mark the task as finished
		defer wg.Done()

		// Simulate some work
		fmt.Println("Task 1: Collect 100 Alien eggs - (0/100)")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 2: Mine 1000 magnite - (0/1000)")
	}()

	// Launch the second task
	go func() {
		// Defer the Done method to mark the task as finished
		defer wg.Done()

		// Simulate some work
		fmt.Println("Task 1: Starting work...")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 1: Work complete!")
	}()

	// Wait for all task to finish
	wg.Wait()

	// All tasks have finished
	fmt.Println("All Tasks completed!")
}
