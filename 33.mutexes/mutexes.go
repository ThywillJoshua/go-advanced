package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateRandomNumber(maxNumber int) int {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	return random.Intn(maxNumber) + 1 // From 1 to maxNumber
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	time.Sleep(time.Duration(generateRandomNumber(5)) * time.Second)
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count++
	fmt.Printf("iteration: %v\n", iteration)
}

func main() {
	var wg sync.WaitGroup

	hitCounter := Hits{}
	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Println("Wait for go routines")
	wg.Wait()

	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()

	fmt.Println("Total:", totalHits)
}
