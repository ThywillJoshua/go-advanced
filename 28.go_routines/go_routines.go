package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}

	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
	}

	for _, char := range data {
		go capIt(char)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("capitalized: %c\n", capitalized)
}
