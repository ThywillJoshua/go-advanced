package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

func generateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	randInt := generateRandInt(1, 100)
	fmt.Println("generateRandInt infinite")

	fmt.Printf("randInt 1: %v\n", <-randInt)
	fmt.Printf("randInt 2: %v\n", <-randInt)
	fmt.Printf("randInt 3: %v\n", <-randInt)

	randIntRange := generateRandIntn(2, 1, 10)

	fmt.Println("generateRandIntN infinite")

	for i := range randIntRange {
		fmt.Printf("Range: %v\n", i)
	}

	randIntn := generateRandIntn(3, 1, 10)
	for {
		n, open := <-randIntn
		if !open {
			fmt.Println("No more numbers")
			break
		}
		fmt.Printf("n: %v\n", n)
	}
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
