package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Divide:
		return lhs / rhs
	case Multiply:
		return lhs * rhs
	}
	panic("Unhandled operation")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))

	subtract := Operation(Subtract)
	fmt.Println(subtract.calculate(4, 2))

	multiply := Operation(Multiply)
	fmt.Println(multiply.calculate(3, 3))

	divide := Operation(Divide)
	fmt.Println(divide.calculate(10, 2))
}
