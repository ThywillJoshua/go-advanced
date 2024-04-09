package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(maxNumber int) int {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	return random.Intn(maxNumber) + 1 // From 1 to maxNumber
}

func main() {
	diceCount := 2
	rollsCount := 8
	sidesCount := 6

	rollSum := 0

	for i := 0; i < rollsCount; i++ {
		for j := 0; j < diceCount; j++ {
			rollSum += generateRandomNumber(sidesCount)
		}
	}

	fmt.Println("Total: ", rollSum)

	if diceCount == 2 && rollSum == 2 {
		fmt.Println("Snake Eyes!")
	}

	if rollSum == 7 {
		fmt.Println("Lucky 7")
	}

	if rollSum%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
