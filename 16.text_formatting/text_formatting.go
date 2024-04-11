package main

import "fmt"

func surround(msg string, left, right rune) string {
	return fmt.Sprintf("%c%v%c", left, msg, right)
}

func main() {
	surrounded := surround("This message", '(', ')')
	fmt.Println(surrounded)
}
