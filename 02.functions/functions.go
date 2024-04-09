package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello ", name)
}

func message() string {
	return "Talk to Jeff"
}

func addNumbers(a, b, c int) int {
	return a + b + c
}

func getTwoNumbers() (int, int) {
	return 2, 3
}

func main() {
	greet("Joshua")

	fmt.Println(message())

	number := addNumbers(1, 2, 3)
	fmt.Println(number)

	one, two := getTwoNumbers()
	fmt.Println(one, two)
}
