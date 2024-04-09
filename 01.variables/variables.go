package main

import "fmt"

func main() {
	var myName string = "Joshua"
	fmt.Println("my name is", myName)

	username := "Thywill Joshua"
	fmt.Println("Username: ", username)

	var sum int
	fmt.Println("Sum: ", sum)

	part1, other := 1, 5
	fmt.Println("Part 1: ", part1, "Other: ", other)

	part2, other := 2, 6
	fmt.Println("Part 2: ", part2, "Other: ", other)

	text1, text2, _ := "Hello", "World", "Hi"
	fmt.Println("text1: ", text1, "text2: ", text2)
}
