package main

import (
	"context"
	"fmt"
	"time"
)

type Vehicle interface {
	calculateSpeed() int
	feature() string
}

// Here Car is concrete class implementing Vehicle interface
func (c Car) calculateSpeed() int {
	fmt.Print(c)
	return 120
}
func (c Car) feature() string {
	return "ZUUUUU"
}

type Car struct {
	tires int
	name  string
}

// only accepts interface, so struct that implements all the methods they are only allowed.
func common(v Vehicle) {
	fmt.Println(v.calculateSpeed())
}

func main() {

	// Create a context with a 20-second timeout

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the cancel function is called to free resources

	// Call a function that simulates a network request
	resultChan := make(chan string)
	go func() {
		// Simulate a long-running operation (e.g., a network request)
		time.Sleep(6 * time.Second) // Simulate work
		resultChan <- "Success!"    // Send result back to the channel
	}()

	// Wait for the result or the context to be canceled
	select {
	case <-ctx.Done():
		// Context was canceled (timeout or manual cancellation)
		fmt.Println("Operation timed out or canceled:", ctx.Err())
	case result := <-resultChan:
		// Operation completed successfully
		fmt.Println("Received result:", result)
	}
	porsh := Car{4, "Porshe"}
	common(porsh)
	// Main()
}
