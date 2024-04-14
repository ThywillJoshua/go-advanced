package main

import "testing"

var player = Player{name: "John", health: 50, energy: 50, maxHealth: 100, maxEnergy: 100}

func TestAddHealth(t *testing.T) {
	player.addHealth(60)
	if player.health > player.maxHealth {
		t.Errorf("Player health should not be more than maxHealth: %v. It's %v", player.maxHealth, player.health)
	}
}

func TestAddEnergy(t *testing.T) {
	player.addEnergy(60)
	if player.energy > player.maxEnergy {
		t.Errorf("Player energy should not be more than maxEnergy: %v. It's %v", player.maxEnergy, player.energy)
	}
}

func TestSubtractHealth(t *testing.T) {
	player.subtractHealth(60)
	if player.health < player.minHealth {
		t.Errorf("Player health should not be more less maxHealth: %v. It's %v", player.minHealth, player.health)
	}
}

func TestSubtractEnergy(t *testing.T) {
	player.subtractEnergy(60)
	if player.energy < player.minEnergy {
		t.Errorf("Player energy should not be less than minEnergy: %v. It's %v", player.minEnergy, player.energy)
	}
}
