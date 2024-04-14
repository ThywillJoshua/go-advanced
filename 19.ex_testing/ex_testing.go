package main

import (
	"fmt"
)

type Player struct {
	name                         string
	health, maxHealth, minHealth uint
	energy, maxEnergy, minEnergy uint
}

func (p *Player) addHealth(amount uint) {
	if (p.health + amount) > p.maxHealth {
		p.health = p.maxHealth
		return
	}
	p.health += amount
}

func (p *Player) subtractHealth(amount uint) {
	if (p.health - amount) < p.minHealth {
		p.health = p.minHealth
		return
	}

	p.health -= amount
}

func (p *Player) addEnergy(amount uint) {
	if (p.energy + amount) > p.maxEnergy {
		p.energy = p.maxEnergy
		return
	}
	p.energy += amount
}

func (p *Player) subtractEnergy(amount uint) {
	if (p.energy - amount) < p.minEnergy {
		p.energy = p.minEnergy
		return
	}

	p.energy -= amount
}

func main() {
	player := Player{name: "John", health: 50, energy: 50, maxHealth: 100, maxEnergy: 100}

	fmt.Println(player)

	player.addHealth(40)
	fmt.Println(player)

	player.addHealth(200)
	fmt.Println(player)
}
