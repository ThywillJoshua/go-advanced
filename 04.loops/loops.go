package main

import "fmt"

func fizzBuzz(number int) {
	if number%5 == 0 && number%3 == 0 {
		fmt.Println("FizzBuzz")
		return
	}

	if number%3 == 0 {
		fmt.Println("Fizz")
		return
	}

	if number%5 == 0 {
		fmt.Println("Buzz")
		return
	}

	fmt.Println(number)
}

func main() {
	number := 1

	for number <= 50 {
		fizzBuzz(number)
		number++
	}
}
