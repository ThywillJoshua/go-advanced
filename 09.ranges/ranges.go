package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, word := range slice {
		fmt.Println(i, word)

		for _, character := range word {
			fmt.Printf("  %q\n", character)
		}
	}
}
