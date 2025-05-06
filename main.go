package main

import (
	"fmt"
	"sync"
)

// Define a struct to represent a point in 2D space
type Point struct {
	X, Y int
}

// Method to move the point by a given offset
func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Procedural function to calculate the sum of two integers
func add(a, b int) int {
	return a + b
}

// Concurrent function to perform a task and send the result to a channel
func concurrentTask(resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Perform some computation
	result := add(10, 20)
	// Send the result to the channel
	resultChan <- result
}

func main() {
	// Create a point and move it
	p := Point{1, 2}
	p.Move(3, 4)
	fmt.Printf("Point: %+v\n", p)

	// Create a channel to receive results from concurrent tasks
	resultChan := make(chan int, 10)

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start multiple concurrent tasks
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go concurrentTask(resultChan, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(resultChan)

	// Collect and print results from the channel
	for result := range resultChan {
		fmt.Printf("Concurrent Task Result: %d\n", result)
	}
}
