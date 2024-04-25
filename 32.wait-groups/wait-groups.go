package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter++
		go goroutine(&wg, &counter)
	}

	fmt.Println("Waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("Done!")
}

func goroutine(wg *sync.WaitGroup, counter *int) {
	defer func() {
		*counter--
		fmt.Println(*counter, "goroutines remaining")
		wg.Done()
	}()

	duration := time.Duration(rand.Intn(5)+1) * time.Second
	fmt.Printf("Waiting for: %v\n", duration)
	time.Sleep(duration)
}
