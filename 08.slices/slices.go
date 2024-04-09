package main

import "fmt"

type Part string

func main() {
	assemblyLine := []Part{
		"Iron", "Battery", "Oil",
	}

	fmt.Println(assemblyLine)

	assemblyLine = append(assemblyLine, "Fuel")
	assemblyLine = append(assemblyLine, "Rubber")
	fmt.Println(assemblyLine)

	assemblyLine = assemblyLine[3:]
	fmt.Println(assemblyLine)
}
