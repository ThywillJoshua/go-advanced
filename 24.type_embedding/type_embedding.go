package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

var sizes = []string{"Small", "Medium", "Large"}

const (
	Ground = iota
	Air
)

var shippingOptions = []string{"Ground", "Air"}

type (
	BeltSize int
	Shipping int
)

func (b BeltSize) String() string {
	return sizes[b]
}

func (s Shipping) String() string {
	return shippingOptions[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomater interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam Mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomater) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)
}
