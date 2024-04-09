package main

import "fmt"

type Coords struct {
	x, y int
}

type Rectangle struct {
	top_left     Coords
	bottom_right Coords
}

func width(rect Rectangle) int {
	return rect.bottom_right.x - rect.top_left.x
}

func lenght(rect Rectangle) int {
	return rect.top_left.y - rect.bottom_right.y
}

func area(rect Rectangle) int {
	return lenght(rect) * width(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) * 2) + (lenght(rect) * 2)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area: ", area(rect))
	fmt.Println("Perimeter: ", perimeter(rect))
}

func main() {
	rect := Rectangle{
		top_left:     Coords{0, 7},
		bottom_right: Coords{10, 0},
	}

	printInfo(rect)

	rect.top_left.y *= 2
	rect.bottom_right.x *= 2

	printInfo(rect)
}
