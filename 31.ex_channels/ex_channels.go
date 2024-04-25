package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job int

func longCalculation(j Job) int {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)

	fmt.Printf("Job %d complete in %v\n", j, duration)

	return int(j) * 30
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 100)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, Job(rand.Intn(10000)))
	}

	return jobs
}

func runJob(resultChan chan int, j Job) {
	resultChan <- longCalculation(j)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	jobs := makeJobs()

	resultsChan := make(chan int, 10)

	for _, job := range jobs {
		go runJob(resultsChan, job)
	}

	resultCount := 0
	sum := 0

	for {
		result := <-resultsChan
		sum += result
		resultCount += 1
		if resultCount == len(jobs) {
			break
		}
	}

	fmt.Println("sum is", sum)
}
