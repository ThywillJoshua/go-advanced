package main

import (
	"fmt"
	"time"
)

// ControlMsg represents control messages for communication between goroutines.
type ControlMsg int

const (
	// DoExit is a control message to indicate exiting the goroutine.
	DoExit = iota
	// ExitOk is a control message to acknowledge successful exit.
	ExitOk
)

// Job represents a job to be processed.
type Job struct {
	data int
}

// Result represents the result of processing a job.
type Result struct {
	result int
	job    Job
}

// doubler is a goroutine that doubles the data of the received jobs.
func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit go routine")
				control <- ExitOk // Acknowledge successful exit
				return
			default:
				panic("Unhandled control message")
			}

		case job := <-jobs:
			// Double the data of the job and send the result to the results channel
			results <- Result{job.data * 2, job}
		}
	}
}

func main() {
	// Create channels for communication between goroutines
	jobs := make(chan Job, 50)       // Channel for sending jobs
	results := make(chan Result, 50) // Channel for receiving results
	control := make(chan ControlMsg) // Channel for sending control messages

	// Start the doubler goroutine
	go doubler(jobs, results, control)

	// Send 30 jobs to the jobs channel
	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			// Process the received result
			fmt.Printf("result: %v\n", result)

		case <-time.After(5 * time.Second):
			// If no result is received within 5 seconds, exit the program
			fmt.Println("timed out")
			control <- DoExit // Send the DoExit control message
			<-control         // Wait for acknowledgement from the doubler goroutine
			fmt.Println("Program Exit")
			return
		}
	}
}
