//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"


type player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

func (p *player) statistic()  {
	fmt.Println("Player:", p.name)
	fmt.Println("Health:", p.health, "/", p.maxHealth)
	fmt.Println("Energy:", p.energy, "/", p.maxEnergy)
}

func (p *player) damageTaken(damage int)  {
	p.health -= damage
	if p.health < 0 {
		p.health = 0
	}
	fmt.Printf("%s took %d damage! Current health: %d/%d\n", p.name, damage, p.health, p.maxHealth)
}

func (p *player) energyUsed(energy int)  {
	p.energy -= energy
	if p.energy < 0 {
		p.energy = 0
	}
	fmt.Printf("%s used %d energy! Current energy: %d/%d\n", p.name, energy, p.energy, p.maxEnergy)
}


func main() {

	player1 := player{
		health:    100,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 50,
		name:      "Masud",
	}
	player1.statistic()
	player1.damageTaken(30)
	player1.energyUsed(20)
	player1.statistic()
}
