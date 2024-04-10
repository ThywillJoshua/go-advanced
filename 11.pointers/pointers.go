package main

import "fmt"

type Tag bool

const (
	Active   = true
	Inactive = false
)

type Product struct {
	name string
	tag  Tag
}

func activate(p *Product) {
	p.tag = Active
}

func deactivate(p *Product) {
	p.tag = Inactive
}

func checkout(products []Product) {
	for i := range products {
		deactivate(&products[i])
	}
}

func main() {
	items := []Product{
		{"Milk", Active}, {"Shirt", Active}, {"Shoe", Active}, {"Socks", Active},
	}

	deactivate(&items[3])
	fmt.Println(items)

	checkout(items)
	fmt.Println(items)
}
