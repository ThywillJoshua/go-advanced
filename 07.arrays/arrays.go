package main

import "fmt"

type Item struct {
	name  string
	price int
}

func printInfo(shoppingList [4]Item) {
	var sum int
	var total int
	var lastItem Item

	for i := 0; i < len(shoppingList); i++ {
		item := shoppingList[i]
		sum += item.price

		if item.name != "" {
			lastItem = item
			total += 1
		}
	}

	fmt.Println("Last item: ", lastItem)
	fmt.Println("Total Items: ", total)
	fmt.Println("Sum: ", sum)
}

func main() {
	shoppingList := [4]Item{
		{"Milk", 200},
		{"Water", 300},
		{"Juice", 100},
	}
	printInfo(shoppingList)

	shoppingList[3] = Item{"House", 50}
	printInfo(shoppingList)
}
